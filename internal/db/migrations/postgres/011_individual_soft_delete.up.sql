-- add created_at column
ALTER TABLE individuals
    ADD COLUMN created_at TIMESTAMP WITH TIME ZONE NULL;

-- add updated_at column
ALTER TABLE individuals
    ADD COLUMN updated_at TIMESTAMP WITH TIME ZONE NULL;

-- add deleted_at column
ALTER TABLE individuals
    ADD COLUMN deleted_at TIMESTAMP WITH TIME ZONE NULL;

-- populate created_at and updated_at with current timestamp
DO
$$
    DECLARE
        t TIMESTAMP WITH TIME ZONE;
    BEGIN
        t := NOW() at time zone 'utc';
        -- noinspection SqlWithoutWhere
        UPDATE individuals
        SET created_at = t,
            updated_at = t;
    END
$$ LANGUAGE plpgsql;

-- set created_at as NOT NULL
ALTER TABLE individuals
    ALTER COLUMN created_at SET NOT NULL;

-- set updated_at as NOT NULL
ALTER TABLE individuals
    ALTER COLUMN updated_at SET NOT NULL;

/*
this procedure is used to prevent updating soft-deleted
records.

this procedure will be triggered before any update
on the individuals table.

Internally, it compares the old and the new record.
If any of the values are being changed, other than the
deleted_at column, it will throw an error.

It relies on the fact that for each update, the application
always sets the updated_at column. So for each update, there
will always be a change detected by this procedure,
and it will successfully throw an error if a user tries
to update a soft-deleted record.
 */
CREATE OR REPLACE FUNCTION prevent_deleted_individual_update()
    RETURNS TRIGGER
AS
$$
BEGIN
    IF (
                TG_OP = 'UPDATE'
            AND OLD.deleted_at IS NOT NULL
            AND EXISTS(
                        SELECT 1
                        FROM jsonb_each(to_jsonb(NEW)) AS post,
                             jsonb_each(to_jsonb(OLD)) AS pre
                        WHERE post.key = pre.key
                          AND (
                            post.key != 'deleted_at' AND post.value IS DISTINCT FROM pre.value
                            )
                          AND post.key != 'deleted_at'
                    )
        ) THEN
        RAISE EXCEPTION 'Cannot update a soft deleted record';

    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- trigger to update updated_at column
CREATE OR REPLACE TRIGGER individual_prevent_deleted_update
    BEFORE UPDATE
    ON individuals
    FOR EACH ROW
EXECUTE PROCEDURE prevent_deleted_individual_update();

