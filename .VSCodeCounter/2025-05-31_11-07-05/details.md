# Details

Date : 2025-05-31 11:07:05

Directory /home/ali/repos/Automation

Total : 153 files,  9072 codes, 347 comments, 1556 blanks, all 10975 lines

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)

## Files
| filename | language | code | comment | blank | total |
| :--- | :--- | ---: | ---: | ---: | ---: |
| [Makefile](/Makefile) | Makefile | 22 | 7 | 8 | 37 |
| [backend.dockerfile](/backend.dockerfile) | Docker | 6 | 0 | 6 | 12 |
| [cmd/mainservice/main.go](/cmd/mainservice/main.go) | Go | 89 | 6 | 7 | 102 |
| [docker-compose.yml](/docker-compose.yml) | YAML | 60 | 0 | 2 | 62 |
| [go.mod](/go.mod) | Go Module File | 49 | 0 | 4 | 53 |
| [go.sum](/go.sum) | Go Checksum File | 147 | 0 | 1 | 148 |
| [internal/api/api/admin.go](/internal/api/api/admin.go) | Go | 203 | 0 | 12 | 215 |
| [internal/api/api/contactinfo.go](/internal/api/api/contactinfo.go) | Go | 132 | 0 | 10 | 142 |
| [internal/api/api/costumapi.go](/internal/api/api/costumapi.go) | Go | 274 | 16 | 33 | 323 |
| [internal/api/api/credentials.go](/internal/api/api/credentials.go) | Go | 188 | 0 | 13 | 201 |
| [internal/api/api/education.go](/internal/api/api/education.go) | Go | 120 | 0 | 9 | 129 |
| [internal/api/api/familyinfo.go](/internal/api/api/familyinfo.go) | Go | 101 | 0 | 8 | 109 |
| [internal/api/api/hospitaldispatch.go](/internal/api/api/hospitaldispatch.go) | Go | 123 | 0 | 12 | 135 |
| [internal/api/api/medicalprofile.go](/internal/api/api/medicalprofile.go) | Go | 147 | 0 | 10 | 157 |
| [internal/api/api/medicine.go](/internal/api/api/medicine.go) | Go | 128 | 0 | 9 | 137 |
| [internal/api/api/militarydetails.go](/internal/api/api/militarydetails.go) | Go | 102 | 0 | 11 | 113 |
| [internal/api/api/person.go](/internal/api/api/person.go) | Go | 204 | 2 | 9 | 215 |
| [internal/api/api/physicalinfo.go](/internal/api/api/physicalinfo.go) | Go | 106 | 0 | 8 | 114 |
| [internal/api/api/prescription.go](/internal/api/api/prescription.go) | Go | 130 | 0 | 9 | 139 |
| [internal/api/api/psychologicalstatus.go](/internal/api/api/psychologicalstatus.go) | Go | 103 | 0 | 8 | 111 |
| [internal/api/api/role.go](/internal/api/api/role.go) | Go | 109 | 0 | 9 | 118 |
| [internal/api/api/service.go](/internal/api/api/service.go) | Go | 99 | 3 | 9 | 111 |
| [internal/api/api/skills.go](/internal/api/api/skills.go) | Go | 116 | 0 | 9 | 125 |
| [internal/config/env.go](/internal/config/env.go) | Go | 48 | 0 | 6 | 54 |
| [internal/core/action/model.go](/internal/core/action/model.go) | Go | 10 | 0 | 3 | 13 |
| [internal/core/action/repo.go](/internal/core/action/repo.go) | Go | 33 | 0 | 8 | 41 |
| [internal/core/action/service.go](/internal/core/action/service.go) | Go | 45 | 2 | 10 | 57 |
| [internal/core/actiontype/model.go](/internal/core/actiontype/model.go) | Go | 11 | 0 | 3 | 14 |
| [internal/core/actiontype/repo.go](/internal/core/actiontype/repo.go) | Go | 33 | 0 | 8 | 41 |
| [internal/core/actiontype/seed.go](/internal/core/actiontype/seed.go) | Go | 38 | 10 | 7 | 55 |
| [internal/core/actiontype/service.go](/internal/core/actiontype/service.go) | Go | 46 | 3 | 10 | 59 |
| [internal/core/admin/model.go](/internal/core/admin/model.go) | Go | 20 | 0 | 5 | 25 |
| [internal/core/admin/repo.go](/internal/core/admin/repo.go) | Go | 46 | 0 | 11 | 57 |
| [internal/core/admin/seed.go](/internal/core/admin/seed.go) | Go | 86 | 1 | 10 | 97 |
| [internal/core/admin/service.go](/internal/core/admin/service.go) | Go | 123 | 7 | 26 | 156 |
| [internal/core/audit/interface.go](/internal/core/audit/interface.go) | Go | 4 | 0 | 2 | 6 |
| [internal/core/bloodgroup/model.go](/internal/core/bloodgroup/model.go) | Go | 5 | 0 | 2 | 7 |
| [internal/core/bloodgroup/repo.go](/internal/core/bloodgroup/repo.go) | Go | 47 | 0 | 10 | 57 |
| [internal/core/bloodgroup/seed.go](/internal/core/bloodgroup/seed.go) | Go | 39 | 1 | 7 | 47 |
| [internal/core/bloodgroup/service.go](/internal/core/bloodgroup/service.go) | Go | 54 | 3 | 11 | 68 |
| [internal/core/contactinfo/model.go](/internal/core/contactinfo/model.go) | Go | 11 | 0 | 2 | 13 |
| [internal/core/contactinfo/repo.go](/internal/core/contactinfo/repo.go) | Go | 45 | 0 | 12 | 57 |
| [internal/core/contactinfo/seed.go](/internal/core/contactinfo/seed.go) | Go | 65 | 1 | 8 | 74 |
| [internal/core/contactinfo/service.go](/internal/core/contactinfo/service.go) | Go | 78 | 5 | 19 | 102 |
| [internal/core/credentials/model.go](/internal/core/credentials/model.go) | Go | 16 | 0 | 4 | 20 |
| [internal/core/credentials/repo.go](/internal/core/credentials/repo.go) | Go | 62 | 0 | 16 | 78 |
| [internal/core/credentials/seed.go](/internal/core/credentials/seed.go) | Go | 66 | 1 | 9 | 76 |
| [internal/core/credentials/service.go](/internal/core/credentials/service.go) | Go | 141 | 5 | 27 | 173 |
| [internal/core/education/model.go](/internal/core/education/model.go) | Go | 12 | 0 | 3 | 15 |
| [internal/core/education/repo.go](/internal/core/education/repo.go) | Go | 45 | 0 | 12 | 57 |
| [internal/core/education/seed.go](/internal/core/education/seed.go) | Go | 63 | 2 | 9 | 74 |
| [internal/core/education/service.go](/internal/core/education/service.go) | Go | 75 | 4 | 11 | 90 |
| [internal/core/educationLevel/model.go](/internal/core/educationLevel/model.go) | Go | 5 | 0 | 2 | 7 |
| [internal/core/educationLevel/repo.go](/internal/core/educationLevel/repo.go) | Go | 38 | 0 | 11 | 49 |
| [internal/core/educationLevel/seed.go](/internal/core/educationLevel/seed.go) | Go | 22 | 1 | 8 | 31 |
| [internal/core/educationLevel/service.go](/internal/core/educationLevel/service.go) | Go | 21 | 0 | 7 | 28 |
| [internal/core/familyinfo/model.go](/internal/core/familyinfo/model.go) | Go | 9 | 0 | 2 | 11 |
| [internal/core/familyinfo/repo.go](/internal/core/familyinfo/repo.go) | Go | 38 | 0 | 11 | 49 |
| [internal/core/familyinfo/seed.go](/internal/core/familyinfo/seed.go) | Go | 58 | 2 | 9 | 69 |
| [internal/core/familyinfo/service.go](/internal/core/familyinfo/service.go) | Go | 65 | 3 | 10 | 78 |
| [internal/core/gender/model.go](/internal/core/gender/model.go) | Go | 5 | 0 | 1 | 6 |
| [internal/core/gender/repo.go](/internal/core/gender/repo.go) | Go | 40 | 0 | 10 | 50 |
| [internal/core/gender/seed.go](/internal/core/gender/seed.go) | Go | 39 | 7 | 7 | 53 |
| [internal/core/gender/service.go](/internal/core/gender/service.go) | Go | 54 | 3 | 11 | 68 |
| [internal/core/hospitaldispatch/model.go](/internal/core/hospitaldispatch/model.go) | Go | 13 | 0 | 4 | 17 |
| [internal/core/hospitaldispatch/repo.go](/internal/core/hospitaldispatch/repo.go) | Go | 38 | 0 | 11 | 49 |
| [internal/core/hospitaldispatch/seed.go](/internal/core/hospitaldispatch/seed.go) | Go | 39 | 1 | 7 | 47 |
| [internal/core/hospitaldispatch/service.go](/internal/core/hospitaldispatch/service.go) | Go | 67 | 4 | 18 | 89 |
| [internal/core/hospitalvisit/model.go](/internal/core/hospitalvisit/model.go) | Go | 17 | 0 | 4 | 21 |
| [internal/core/hospitalvisit/repo.go](/internal/core/hospitalvisit/repo.go) | Go | 38 | 0 | 11 | 49 |
| [internal/core/hospitalvisit/seed.go](/internal/core/hospitalvisit/seed.go) | Go | 30 | 1 | 6 | 37 |
| [internal/core/hospitalvisit/service.go](/internal/core/hospitalvisit/service.go) | Go | 57 | 4 | 11 | 72 |
| [internal/core/medicalprofile/model.go](/internal/core/medicalprofile/model.go) | Go | 20 | 0 | 4 | 24 |
| [internal/core/medicalprofile/repo.go](/internal/core/medicalprofile/repo.go) | Go | 48 | 0 | 12 | 60 |
| [internal/core/medicalprofile/seed.go](/internal/core/medicalprofile/seed.go) | Go | 41 | 1 | 7 | 49 |
| [internal/core/medicalprofile/service.go](/internal/core/medicalprofile/service.go) | Go | 70 | 4 | 19 | 93 |
| [internal/core/medicines/model.go](/internal/core/medicines/model.go) | Go | 8 | 0 | 2 | 10 |
| [internal/core/medicines/repo.go](/internal/core/medicines/repo.go) | Go | 38 | 0 | 9 | 47 |
| [internal/core/medicines/seed.go](/internal/core/medicines/seed.go) | Go | 44 | 1 | 5 | 50 |
| [internal/core/medicines/service.go](/internal/core/medicines/service.go) | Go | 67 | 4 | 18 | 89 |
| [internal/core/militarydetails/model.go](/internal/core/militarydetails/model.go) | Go | 13 | 0 | 3 | 16 |
| [internal/core/militarydetails/repo.go](/internal/core/militarydetails/repo.go) | Go | 38 | 0 | 11 | 49 |
| [internal/core/militarydetails/seed.go](/internal/core/militarydetails/seed.go) | Go | 75 | 1 | 9 | 85 |
| [internal/core/militarydetails/service.go](/internal/core/militarydetails/service.go) | Go | 75 | 4 | 10 | 89 |
| [internal/core/person/model.go](/internal/core/person/model.go) | Go | 64 | 1 | 11 | 76 |
| [internal/core/person/repo.go](/internal/core/person/repo.go) | Go | 91 | 0 | 15 | 106 |
| [internal/core/person/seed.go](/internal/core/person/seed.go) | Go | 146 | 2 | 17 | 165 |
| [internal/core/person/service.go](/internal/core/person/service.go) | Go | 153 | 11 | 27 | 191 |
| [internal/core/persontype/model.go](/internal/core/persontype/model.go) | Go | 5 | 0 | 2 | 7 |
| [internal/core/persontype/repo.go](/internal/core/persontype/repo.go) | Go | 40 | 0 | 9 | 49 |
| [internal/core/persontype/seed.go](/internal/core/persontype/seed.go) | Go | 36 | 13 | 8 | 57 |
| [internal/core/persontype/service.go](/internal/core/persontype/service.go) | Go | 54 | 3 | 9 | 66 |
| [internal/core/physicalinfo/model.go](/internal/core/physicalinfo/model.go) | Go | 20 | 0 | 4 | 24 |
| [internal/core/physicalinfo/repo.go](/internal/core/physicalinfo/repo.go) | Go | 40 | 0 | 11 | 51 |
| [internal/core/physicalinfo/seed.go](/internal/core/physicalinfo/seed.go) | Go | 82 | 8 | 12 | 102 |
| [internal/core/physicalinfo/service.go](/internal/core/physicalinfo/service.go) | Go | 68 | 3 | 10 | 81 |
| [internal/core/physicalstatus/model.go](/internal/core/physicalstatus/model.go) | Go | 5 | 0 | 2 | 7 |
| [internal/core/physicalstatus/repo.go](/internal/core/physicalstatus/repo.go) | Go | 34 | 0 | 7 | 41 |
| [internal/core/physicalstatus/seed.go](/internal/core/physicalstatus/seed.go) | Go | 31 | 1 | 7 | 39 |
| [internal/core/physicalstatus/service.go](/internal/core/physicalstatus/service.go) | Go | 35 | 1 | 9 | 45 |
| [internal/core/prescription/model.go](/internal/core/prescription/model.go) | Go | 15 | 0 | 4 | 19 |
| [internal/core/prescription/repo.go](/internal/core/prescription/repo.go) | Go | 38 | 0 | 9 | 47 |
| [internal/core/prescription/seed.go](/internal/core/prescription/seed.go) | Go | 39 | 1 | 7 | 47 |
| [internal/core/prescription/service.go](/internal/core/prescription/service.go) | Go | 67 | 4 | 18 | 89 |
| [internal/core/psychologicalstatus/model.go](/internal/core/psychologicalstatus/model.go) | Go | 5 | 0 | 3 | 8 |
| [internal/core/psychologicalstatus/repo.go](/internal/core/psychologicalstatus/repo.go) | Go | 33 | 0 | 9 | 42 |
| [internal/core/psychologicalstatus/seed.go](/internal/core/psychologicalstatus/seed.go) | Go | 31 | 1 | 7 | 39 |
| [internal/core/psychologicalstatus/service.go](/internal/core/psychologicalstatus/service.go) | Go | 50 | 3 | 14 | 67 |
| [internal/core/rank/model.go](/internal/core/rank/model.go) | Go | 6 | 0 | 1 | 7 |
| [internal/core/rank/repo.go](/internal/core/rank/repo.go) | Go | 41 | 0 | 11 | 52 |
| [internal/core/rank/seed.go](/internal/core/rank/seed.go) | Go | 67 | 1 | 7 | 75 |
| [internal/core/rank/service.go](/internal/core/rank/service.go) | Go | 78 | 3 | 11 | 92 |
| [internal/core/religion/model.go](/internal/core/religion/model.go) | Go | 6 | 0 | 1 | 7 |
| [internal/core/religion/repo.go](/internal/core/religion/repo.go) | Go | 40 | 0 | 9 | 49 |
| [internal/core/religion/seed.go](/internal/core/religion/seed.go) | Go | 54 | 15 | 7 | 76 |
| [internal/core/religion/service.go](/internal/core/religion/service.go) | Go | 58 | 3 | 11 | 72 |
| [internal/core/role/model.go](/internal/core/role/model.go) | Go | 6 | 0 | 1 | 7 |
| [internal/core/role/repo.go](/internal/core/role/repo.go) | Go | 45 | 0 | 12 | 57 |
| [internal/core/role/seed.go](/internal/core/role/seed.go) | Go | 47 | 1 | 7 | 55 |
| [internal/core/role/service.go](/internal/core/role/service.go) | Go | 64 | 3 | 11 | 78 |
| [internal/core/skills/model.go](/internal/core/skills/model.go) | Go | 11 | 0 | 3 | 14 |
| [internal/core/skills/repo.go](/internal/core/skills/repo.go) | Go | 45 | 0 | 12 | 57 |
| [internal/core/skills/seed.go](/internal/core/skills/seed.go) | Go | 68 | 3 | 9 | 80 |
| [internal/core/skills/serivce.go](/internal/core/skills/serivce.go) | Go | 68 | 3 | 11 | 82 |
| [internal/db/db.go](/internal/db/db.go) | Go | 38 | 0 | 8 | 46 |
| [internal/db/migrate.go](/internal/db/migrate.go) | Go | 57 | 1 | 3 | 61 |
| [internal/logger/logger.go](/internal/logger/logger.go) | Go | 24 | 0 | 6 | 30 |
| [internal/logger/logger.md](/internal/logger/logger.md) | Markdown | 51 | 0 | 34 | 85 |
| [internal/middleware/auth.go](/internal/middleware/auth.go) | Go | 86 | 0 | 5 | 91 |
| [internal/middleware/headers.go](/internal/middleware/headers.go) | Go | 39 | 0 | 8 | 47 |
| [internal/middleware/middleware.md](/internal/middleware/middleware.md) | Markdown | 65 | 0 | 42 | 107 |
| [internal/response/response.go](/internal/response/response.go) | Go | 34 | 0 | 8 | 42 |
| [internal/seeder/seed.go](/internal/seeder/seed.go) | Go | 92 | 2 | 9 | 103 |
| [migrations/000001\_create\_tables.down.sql](/migrations/000001_create_tables.down.sql) | MS SQL | 17 | 0 | 0 | 17 |
| [migrations/000001\_create\_tables.up.sql](/migrations/000001_create_tables.up.sql) | MS SQL | 129 | 0 | 16 | 145 |
| [pkg/security/jwt.go](/pkg/security/jwt.go) | Go | 39 | 1 | 7 | 47 |
| [pkg/security/password.go](/pkg/security/password.go) | Go | 10 | 1 | 4 | 15 |
| [pkg/services/router/router.go](/pkg/services/router/router.go) | Go | 10 | 1 | 4 | 15 |
| [pkg/services/router/routes.go](/pkg/services/router/routes.go) | Go | 89 | 24 | 10 | 123 |
| [readme.md](/readme.md) | Markdown | 399 | 0 | 56 | 455 |
| [test (2)/http/databseservice/admin.http](/test%20(2)/http/databseservice/admin.http) | HTTP | 87 | 16 | 25 | 128 |
| [test (2)/http/databseservice/contactinfo.http](/test%20(2)/http/databseservice/contactinfo.http) | HTTP | 41 | 11 | 24 | 76 |
| [test (2)/http/databseservice/createperson.http](/test%20(2)/http/databseservice/createperson.http) | HTTP | 68 | 5 | 15 | 88 |
| [test (2)/http/databseservice/credentials.http](/test%20(2)/http/databseservice/credentials.http) | HTTP | 59 | 15 | 38 | 112 |
| [test (2)/http/databseservice/education.http](/test%20(2)/http/databseservice/education.http) | HTTP | 33 | 8 | 20 | 61 |
| [test (2)/http/databseservice/familyinfo.http](/test%20(2)/http/databseservice/familyinfo.http) | HTTP | 28 | 6 | 16 | 50 |
| [test (2)/http/databseservice/htttp.http](/test%20(2)/http/databseservice/htttp.http) | HTTP | 88 | 10 | 20 | 118 |
| [test (2)/http/databseservice/militarydetails.http](/test%20(2)/http/databseservice/militarydetails.http) | HTTP | 30 | 5 | 7 | 42 |
| [test (2)/http/databseservice/person.http](/test%20(2)/http/databseservice/person.http) | HTTP | 75 | 11 | 17 | 103 |
| [test (2)/http/databseservice/physicalinfo.http](/test%20(2)/http/databseservice/physicalinfo.http) | HTTP | 39 | 6 | 13 | 58 |
| [test (2)/http/databseservice/physicalstatus.http](/test%20(2)/http/databseservice/physicalstatus.http) | HTTP | 33 | 6 | 9 | 48 |
| [test (2)/http/databseservice/role.http](/test%20(2)/http/databseservice/role.http) | HTTP | 33 | 7 | 10 | 50 |
| [test (2)/http/databseservice/skills.http](/test%20(2)/http/databseservice/skills.http) | HTTP | 38 | 7 | 10 | 55 |

[Summary](results.md) / Details / [Diff Summary](diff.md) / [Diff Details](diff-details.md)