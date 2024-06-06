package db

import (
	"context"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/nrc-no/notcore/internal/api"
	"github.com/nrc-no/notcore/internal/constants"
	"github.com/nrc-no/notcore/internal/containers"
	"github.com/nrc-no/notcore/internal/utils"
	"github.com/nrc-no/notcore/pkg/api/deduplication"
	"github.com/stretchr/testify/assert"
)

func Seed(ctx context.Context, sqlDb *sqlx.DB) *api.Country {
	countryRepo := NewCountryRepo(sqlDb)
	countryDef := &api.Country{
		Code: "NO",
		Name: "Norway",
		ReadGroup: "read",
		WriteGroup: "write",
	}
	country, err := countryRepo.Put(ctx, countryDef)
	if err != nil {
		log.Fatalf("Failed to seed database: %s", err)
	}
	return country
}

func ClearDatabase(ctx context.Context, sqlDb *sqlx.DB) {
	_, err := sqlDb.ExecContext(ctx, "TRUNCATE TABLE individual_registrations CASCADE")
	if err != nil {
		log.Fatalf("Failed to clear database: %s", err)
	}	
}

type TestSpec struct {
	name string
	seed []*api.Individual
	individuals []*api.Individual
	deduplicationConfig deduplication.DeduplicationConfig
	wantFile []containers.Set[int]
	wantDb map[int][]int
}

func TestDeduplication(t *testing.T) {
	pool, resource := InitTestDocker("5432")
	defer pool.Purge(resource)

	ctx := context.Background()
	sqlDb := OpenDatabaseConnection(ctx, pool, resource, "5432")
	defer sqlDb.Close()

	RunMigrations(ctx, sqlDb)

	country := Seed(ctx, sqlDb)
	ctx = utils.WithSelectedCountryID(ctx, country.ID)

	wantFile2Equal := []containers.Set[int]{
		containers.NewSet[int](1),
		containers.NewSet[int](0),
	}

	wantDb2Equal := map[int][]int{
		0: {0},
	}

	basicRulesFindDuplicateInFile := []TestSpec {
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Ids (1=1); [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					IdentificationNumber1: "0987654321",
				},
			},
			individuals: []*api.Individual{
				{
					IdentificationNumber1: "1234567890",
				},
				{
					IdentificationNumber1: "1234567890",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Phone numbers (1=1); [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					PhoneNumber1: "0987654321",
				},
			},
			individuals: []*api.Individual{
				{
					PhoneNumber1: "1234567890",
				},
				{
					PhoneNumber1: "1234567890",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNamePhoneNumbers],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Emails; [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					Email1: "other@not.nrc.no",
				},
			},
			individuals: []*api.Individual{
				{
					Email1: "test@not.nrc.no",
				},
				{
					Email1: "test@not.nrc.no",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameEmails],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Names (all set); [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					FirstName: "Jane",
					MiddleName: "Doe",
					LastName: "Smith",
					NativeName: "Jane Doe",
				},
			},
			individuals: []*api.Individual{
				{
					FirstName: "John",
					MiddleName: "Doe",
					LastName: "Smith",
					NativeName: "John Doe",
				},
				{
					FirstName: "John",
					MiddleName: "Doe",
					LastName: "Smith",
					NativeName: "John Doe",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FullName; [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					FullName: "Jane Doe",
				},
			},
			individuals: []*api.Individual{
				{
					FullName: "John Doe",
				},
				{
					FullName: "John Doe",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					 deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField1; [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField1: "Dolor sit amet",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField1: "Lorem ipsum",
				},
				{
					FreeField1: "Lorem ipsum",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField1],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField2; [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField2: "Dolor sit amet",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField2: "Lorem ipsum",
				},
				{
					FreeField2: "Lorem ipsum",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField2],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField3; [FileOrDB] File; [Expected] 1;",
			seed:[]*api.Individual{
				{
					FreeField3: "Dolor sit amet",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField3: "Lorem ipsum",
				},
				{
					FreeField3: "Lorem ipsum",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField3],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField4; [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField4: "Dolor sit amet",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField4: "Lorem ipsum",
				},
				{
					FreeField4: "Lorem ipsum",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField4],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField5; [FileOrDB] File; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField5: "Dolor sit amet",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField5: "Lorem ipsum",
				},
				{
					FreeField5: "Lorem ipsum",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField5],
				},
			},
			wantFile: wantFile2Equal,
			wantDb: nil,
		},
	}

	basicRulesFindDuplicateInDb := []TestSpec{
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Phone numbers (1=1); [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					PhoneNumber1: "1234567890",
				},
			},
			individuals: []*api.Individual{
				{
					PhoneNumber1: "1234567890",
				},
				{
					PhoneNumber1: "0987654321",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNamePhoneNumbers],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Emails; [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					Email1: "test@not.nrc.no",
				},
			},
			individuals: []*api.Individual{
				{
					Email1: "test@not.nrc.no",
				},
				{
					Email1: "other@not.nrc.no",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameEmails],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Names (all set); [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					FirstName:  "John",
					MiddleName: "Doe",
					LastName:   "Smith",
					NativeName: "John Doe",
				},
			},
			individuals: []*api.Individual{
				{
					FirstName:  "John",
					MiddleName: "Doe",
					LastName:   "Smith",
					NativeName: "John Doe",
				},
				{
					FirstName:  "Jane",
					MiddleName: "Doe",
					LastName:   "Smith",
					NativeName: "Jane Doe",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FullName; [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					FullName: "John Doe",
				},
			},
			individuals: []*api.Individual{
				{
					FullName: "John Doe",
				},
				{
					FullName: "Jane Doe",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField1; [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField1: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField1: "Lorem ipsum",
				},
				{
					FreeField1: "Dolor sit amet",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField1],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField2; [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField2: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField2: "Lorem ipsum",
				},
				{
					FreeField2: "Dolor sit amet",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField2],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField3; [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField3: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField3: "Lorem ipsum",
				},
				{
					FreeField3: "Dolor sit amet",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField3],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField4; [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField4: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField4: "Lorem ipsum",
				},
				{
					FreeField4: "Dolor sit amet",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField4],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField5; [FileOrDB] DB; [Expected] 1;",
			seed: []*api.Individual{
				{
					FreeField5: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField5: "Lorem ipsum",
				},
				{
					FreeField5: "Dolor sit amet",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField5],
				},
			},
			wantFile: nil,
			wantDb:   wantDb2Equal,
		},
	}

	basicRulesDontFindDuplicate := []TestSpec{
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Ids; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					IdentificationNumber1: "1234567890",
				},
			},
			individuals: []*api.Individual{
				{
					IdentificationNumber1: "0987654321",
				},
				{
					IdentificationNumber1: "6666666666",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameIds],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Phone numbers; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					PhoneNumber1: "1234567890",
				},
			},
			individuals: []*api.Individual{
				{
					PhoneNumber1: "0987654321",
				},
				{
					PhoneNumber1: "6666666666",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNamePhoneNumbers],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Emails; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					Email1: "a@not.nrc.no",
				},
			},
			individuals: []*api.Individual{
				{
					Email1: "b@not.nrc.no",
				},
				{
					Email1: "c@not.nrc.no",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameEmails],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] Names (all set); [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					FirstName:  "John",
					MiddleName: "Doe",
					LastName:   "Smith",
					NativeName: "John Doe",
				},
			},
			individuals: []*api.Individual{
				{
					FirstName:  "Jane",
					MiddleName: "Doe",
					LastName:   "Smith",
					NativeName: "Jane Doe",
				},
				{
					FirstName:  "James",
					MiddleName: "Doe",
					LastName:   "Smith",
					NativeName: "James Doe",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameNames],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FullName; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					FullName: "John Doe",
				},
			},
			individuals: []*api.Individual{
				{
					FullName: "Jane Doe",
				},
				{
					FullName: "James Doe",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					 deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFullName],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField1; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					FreeField1: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField1: "Dolor sit amet",
				},
				{
					FreeField1: "Consectetur adipiscing elit",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField1],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField2; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					FreeField2: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField2: "Dolor sit amet",
				},
				{
					FreeField2: "Consectetur adipiscing elit",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					 deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField2],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField3; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					FreeField3: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField3: "Dolor sit amet",
				},
				{
					FreeField3: "Consectetur adipiscing elit",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					 deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField3],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField4; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					FreeField4: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField4: "Dolor sit amet",
				},
				{
					FreeField4: "Consectetur adipiscing elit",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					 deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField4],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
		{
			name: "[Duplicate.LogicOperator] AND; [Duplicate.Fields] FreeField5; [FileOrDB] File; [Expected] 0;",
			seed: []*api.Individual{
				{
					FreeField5: "Lorem ipsum",
				},
			},
			individuals: []*api.Individual{
				{
					FreeField5: "Dolor sit amet",
				},
				{
					FreeField5: "Consectetur adipiscing elit",
				},
			},
			deduplicationConfig: deduplication.DeduplicationConfig{
				Operator: deduplication.LOGICAL_OPERATOR_AND,
				Types: []deduplication.DeduplicationType{
					 deduplication.DeduplicationTypes[deduplication.DeduplicationTypeNameFreeField5],
				},
			},
			wantFile: nil,
			wantDb: nil,
		},
	}

	tests := append(
		append(
			basicRulesFindDuplicateInFile,
			basicRulesFindDuplicateInDb...,
		),
		basicRulesDontFindDuplicate...,
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ClearDatabase(ctx, sqlDb)

			individualRepo := NewIndividualRepo(sqlDb)

			for _, ind := range tt.individuals {
				ind.CountryID = country.ID
			}

			for _, ind := range tt.seed {
				ind.CountryID = country.ID
			}

			seedInds, err := individualRepo.PutMany(ctx, tt.seed, constants.IndividualDBColumns)
			if err != nil {
				t.Fatalf("Failed to seed database: %s", err)
			}

			fileDupes, dbDupes, err := individualRepo.FindDuplicates(ctx, tt.individuals, tt.deduplicationConfig)

			if err != nil {
				t.Fatalf("Failed to deduplicate: %s", err)
			}

			wantDbInds := make(map[int][]*api.Individual, 0)
			for i, ind := range tt.wantDb {
				wantDbInds[i] = make([]*api.Individual, 0)
				for _, j := range ind {
					wantDbInds[i] = append(wantDbInds[i], seedInds[j])
				} 
			}

			assert.ElementsMatch(t, tt.wantFile, fileDupes)

			for i, indList := range dbDupes {
				for j, ind := range indList {
					assert.Equal(t, wantDbInds[i][j].ID, ind.ID)
				}
			}
		})
	}
}