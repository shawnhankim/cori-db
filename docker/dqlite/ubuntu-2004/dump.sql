PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE test_table(id int, description varchar(10));
INSERT INTO test_table VALUES(1,'foo');
INSERT INTO test_table VALUES(2,'bar');
COMMIT;
