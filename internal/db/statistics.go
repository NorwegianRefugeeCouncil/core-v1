package db

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
)

type StatisticRepo interface {
	CountryStatistics(ctx context.Context, countryID string) (api.CountryStatistics, error)
}

func NewStatisticsRepo(db *sqlx.DB) StatisticRepo {
	return &statisticsRepo{db: db}
}

type statisticsRepo struct {
	db *sqlx.DB
}

func (i statisticsRepo) CountryStatistics(ctx context.Context, countryID string) (api.CountryStatistics, error) {
	var stats api.CountryStatistics
	stats.ByGenderAge = make(map[string]map[string]int)
	stats.ByDisplacementStatus = make(map[string]int)
	stats.BySensoryImpairment = make(map[string]int)
	stats.ByMentalImpairment = make(map[string]int)
	stats.ByPhysicalImpairment = make(map[string]int)

	type ResultRow struct {
		Gender             *string `db:"gender"`
		DisplacementStatus *string `db:"displacement_status"`
		PhysicalImpairment *string `db:"physical_impairment"`
		MentalImpairment   *string `db:"mental_impairment"`
		SensoryImpairment  *string `db:"sensory_impairment"`
		AgeGroup           *string `db:"age_group"`
		Majority           *string `db:"majority"`
		Count              int     `db:"count"`
	}

	qry := `
WITH source AS (
	SELECT 
	    CASE WHEN is_minor THEN 'minor' 
	         WHEN EXTRACT(YEAR FROM AGE(NOW(), birth_date)) >= 65 THEN 'elderly'
	         ELSE 'adult' END AS majority,
		CASE WHEN gender = '' THEN 'unknown' ELSE gender END AS gender,
		CASE WHEN displacement_status = '' then 'unknown' ELSE displacement_status END AS displacement_status,
		CASE WHEN physical_impairment = '' THEN 'unknown' ELSE physical_impairment END AS physical_impairment,
		CASE WHEN mental_impairment = '' THEN 'unknown' ELSE mental_impairment END AS mental_impairment,
		CASE WHEN sensory_impairment = '' THEN 'unknown' ELSE sensory_impairment END AS sensory_impairment,
		EXTRACT(YEAR FROM AGE(NOW(), birth_date)) as age
	FROM individuals
	WHERE country_id = $1
)

SELECT 
    count(*) as count,
    majority,
	gender,
	displacement_status,
	physical_impairment,
	mental_impairment,
	sensory_impairment,
	CASE
	WHEN age IS NULL THEN 'unknown'
	WHEN age < 18 THEN '0-17'
	WHEN age BETWEEN 18 AND 24 THEN '18-24'
	WHEN age BETWEEN 25 AND 34 THEN '25-34'
	WHEN age BETWEEN 35 AND 44 THEN '35-44'
	WHEN age BETWEEN 45 AND 54 THEN '45-54'
	WHEN age BETWEEN 55 AND 64 THEN '55-64'
	WHEN age BETWEEN 65 AND 74 THEN '65-74'
	WHEN age BETWEEN 75 AND 84 THEN '75-84'
	WHEN age >= 85 THEN '85+'
	END AS age_group
FROM source
GROUP BY GROUPING SETS (
	(),
    (majority),
    (gender, majority),
    (gender, age_group),
    (displacement_status),
    (physical_impairment),
    (mental_impairment),
    (sensory_impairment)
)
`

	var groupedRows []ResultRow
	err := i.db.SelectContext(ctx, &groupedRows, qry, countryID)
	if err != nil {
		return api.CountryStatistics{}, err
	}

	for _, row := range groupedRows {
		if row.Gender != nil && row.AgeGroup != nil {
			gender := *row.Gender
			ageGroup := *row.AgeGroup
			if _, ok := stats.ByGenderAge[gender]; !ok {
				stats.ByGenderAge[gender] = map[string]int{}
			}
			stats.ByGenderAge[gender][ageGroup] = row.Count
		} else if row.DisplacementStatus != nil {
			stats.ByDisplacementStatus[*row.DisplacementStatus] = row.Count
		} else if row.PhysicalImpairment != nil {
			stats.ByPhysicalImpairment[*row.PhysicalImpairment] = row.Count
		} else if row.MentalImpairment != nil {
			stats.ByMentalImpairment[*row.MentalImpairment] = row.Count
		} else if row.SensoryImpairment != nil {
			stats.BySensoryImpairment[*row.SensoryImpairment] = row.Count
		} else if row.Gender != nil && row.Majority != nil {
			gender := *row.Gender
			majority := *row.Majority
			if gender == "female" && majority != "minor" {
				stats.WomenAndChildrenCount += row.Count
			}
			if gender != "female" && majority == "minor" {
				stats.WomenAndChildrenCount += row.Count
			}
			if gender == "female" && majority == "minor" {
				stats.WomenAndChildrenCount += row.Count
			}
		} else if row.Majority != nil {
			if *row.Majority == "minor" {
				stats.ChildrenCount = row.Count
			} else if *row.Majority == "elderly" {
				stats.ElderlyCount = row.Count
			}
		} else {
			stats.TotalCount = row.Count
		}
	}

	return stats, nil
}
