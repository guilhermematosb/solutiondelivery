export PGPASSWORD=admin

# Database DDL
psql -U postgres -d solutiondelivery -a -f /docker-entrypoint-initdb.d/ddl/init.sql
psql -U postgres -d solutiondelivery -a -f /docker-entrypoint-initdb.d/ddl/solutiondelivery.sql

# Test Database DDL
psql -U postgres -d solutiondelivery_test -a -f /docker-entrypoint-initdb.d/ddl/solutiondelivery_test.sql
