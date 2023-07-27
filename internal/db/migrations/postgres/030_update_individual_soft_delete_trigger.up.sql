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
        RAISE EXCEPTION 'Cannot update a soft deleted record %', OLD.id::text;

    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- trigger to update updated_at column
CREATE OR REPLACE TRIGGER individual_registrations_prevent_deleted_update
    BEFORE UPDATE
    ON individual_registrations
    FOR EACH ROW
EXECUTE PROCEDURE prevent_deleted_individual_update();

