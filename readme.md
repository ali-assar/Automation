

# Using the PersonInfo System: A Step-by-Step Guide

This guide provides a comprehensive walkthrough for interacting with the PersonInfo system API. You will learn how to authenticate, retrieve static data, and create a new person record using the provided endpoints. The API runs locally on `localhost:8081` and uses JSON for data exchange.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Step 1: Authenticate with the System (Login)](#step-1-authenticate-with-the-system-login)
3. [Step 2: Retrieve Static Tables](#step-2-retrieve-static-tables)
4. [Step 3: Create a Full Person Record](#step-3-create-a-full-person-record)
5. [Troubleshooting](#troubleshooting)

---

## Prerequisites

Before you begin, ensure you have the following:

- **API Server Running**: The PersonInfo system API should be running on `http://localhost:8081`. If it’s not running, start the server (e.g., using `go run main.go` if it’s a Go application).
- **HTTP Client**: A tool to send HTTP requests, such as:
  - **VS Code REST Client** (using a `.http` file).
  - **Postman**.
  - **cURL**.
- **Admin Credentials**: You’ll need a valid `username` and `password` for authentication. For this guide, we’ll use:
  - Username: `admin1`
  - Password: `admin123`
- **Database Setup**: Ensure the database (e.g., SQLite `test.db`) is set up and contains the necessary static data (e.g., blood groups, religions, person types, ranks).

---

## Step 1: Authenticate with the System (Login)

The first step is to authenticate with the system by logging in. This will provide you with two JWT tokens (`dynamic_token` and `static_token`) required for accessing protected endpoints.

### Request
Send a `POST` request to the `/api/personinfo/login` endpoint with your credentials.

#### Example Request
```http
POST http://localhost:8081/api/personinfo/login
Content-Type: application/json

{
  "username": "admin1",
  "password": "admin123"
}
```

- **URL**: `http://localhost:8081/api/personinfo/login`
- **Method**: `POST`
- **Headers**:
  - `Content-Type: application/json`
- **Body**:
  - `username`: The admin username (e.g., `"admin1"`).
  - `password`: The admin password (e.g., `"admin123"`).

### Expected Response
If the credentials are correct, you’ll receive a `200 OK` response with the admin details and tokens.

#### Example Response
```json
HTTP/1.1 200 OK
Access-Control-Allow-Headers: Content-Type, Authorization, Content-Length, Accept-Encoding, Cache-Control, X-Requested-With, accept, origin
Access-Control-Allow-Methods: GET, POST, PUT, DELETE
Access-Control-Allow-Origin: *
Content-Security-Policy: default-src 'none'
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 May 2025 07:02:58 GMT
Connection: close
Transfer-Encoding: chunked

{
  "admin": {
    "ID": "083308ac-3c6b-4c5a-a1f5-6dfc63d4d63f",
    "NationalIDNumber": "012345678",
    "UserName": "admin1",
    "HashPassword": "$2a$10$uTlUpmNVpGhqq1OXm3SWv.uk5BH4PXYdqH3KUGsTntaEQEUwjw6uK",
    "RoleID": 1,
    "DeletedAt": 0,
    "CredentialsID": 0,
    "Person": {
      "NationalIDNumber": "012345678",
      "FirstName": "John",
      "LastName": "Doe",
      "FamilyInfoID": {
        "Int64": 1,
        "Valid": true
      },
      "ContactInfoID": {
        "Int64": 1,
        "Valid": true
      },
      "SkillsID": {
        "Int64": 1,
        "Valid": true
      },
      "PhysicalInfoID": {
        "Int64": 1,
        "Valid": true
      },
      "BirthDate": "1990-01-01T00:00:00Z",
      "ReligionID": {
        "Int64": 1,
        "Valid": true
      },
      "PersonTypeID": {
        "Int64": 1,
        "Valid": true
      },
      "MilitaryDetailsID": {
        "Int64": 1,
        "Valid": true
      },
      "DeletedAt": 0,
      "FamilyInfo": {
        "ID": 0,
        "FatherDetails": "",
        "MotherDetails": "",
        "ChildsDetails": "",
        "HusbandDetails": "",
        "DeletedAt": 0
      },
      "ContactInfo": {
        "ID": 0,
        "Address": "",
        "PhoneNumber": "",
        "EmergencyPhoneNumber": "",
        "LandlinePhone": "",
        "EmailAddress": "",
        "SocialMedia": "",
        "DeletedAt": 0
      },
      "Skills": {
        "ID": 0,
        "EducationID": 0,
        "Languages": "",
        "SkillsDescription": "",
        "Certificates": "",
        "DeletedAt": 0,
        "Education": {
          "ID": 0,
          "EducationLevelID": 0,
          "FieldOfStudy": 0,
          "Description": "",
          "University": "",
          "StartDate": 0,
          "EndDate": 0,
          "DeletedAt": 0
        }
      },
      "Religion": {
        "ID": 0,
        "ReligionName": "",
        "ReligionType": ""
      },
      "PersonType": {
        "ID": 0,
        "Type": ""
      },
      "MilitaryDetails": {
        "ID": 0,
        "RankID": 0,
        "ServiceStartDate": 0,
        "ServiceDispatchDate": 0,
        "ServiceUnit": 0,
        "BattalionUnit": 0,
        "CompanyUnit": 0,
        "DeletedAt": 0,
        "RankRef": {
          "ID": 0,
          "Name": "",
          "DeletedAt": 0
        }
      },
      "PhysicalInfo": {
        "ID": 0,
        "Height": 0,
        "Weight": 0,
        "EyeColor": "",
        "BloodGroupID": 0,
        "GenderID": 0,
        "PhysicalStatusID": 0,
        "DeletedAt": 0,
        "BloodGroup": {
          "ID": 0,
          "Name": ""
        },
        "Gender": {
          "ID": 0,
          "Gender": ""
        },
        "PhysicalStatus": {
          "ID": 0,
          "Status": "",
          "Description": "",
          "DeletedAt": 0
        }
      }
    },
    "Role": {
      "ID": 0,
      "Type": "",
      "DeletedAt": 0
    }
  },
  "dynamic_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjA4MzMwOGFjLTNjNmItNGM1YS1hMWY1LTZkZmM2M2Q0ZDYzZiIsImV4cCI6MTc0ODQxNjY3OCwicm9sZSI6MX0.MZtQzhzQS_HmE_k9V-l_g5ImHYAgbx311kY_r0XgA5g",
  "static_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjA4MzMwOGFjLTNjNmItNGM1YS1hMWY1LTZkZmM2M2Q0ZDYzZiIsImV4cCI6MTc0ODUwMjE3OCwicm9sZSI6MX0.VjRyV3chUPrqOPiC0GK_aLX9nFkY7uiSbyK8-QP0_ew"
}
```

- **Status**: `200 OK`
- **Key Fields**:
  - `admin`: Contains the admin’s details, including their associated `Person` data.
  - `dynamic_token`: A JWT token for accessing dynamic endpoints (e.g., creating or updating records). It expires at the timestamp `1748416678` (May 28, 2025, 10:17 AM EEST).
  - `static_token`: A JWT token for accessing static endpoints (e.g., retrieving static tables). It expires at `1748502178` (May 29, 2025, 10:02 AM EEST).

### Action
Copy the `static_token` and `dynamic_token` from the response. You’ll need the `static_token` for the next step (retrieving static tables) and the `dynamic_token` for creating a person record.

---

## Step 2: Retrieve Static Tables

The `/api/personinfo/static/static-tables` endpoint provides static data (e.g., blood groups, religions, person types, ranks) needed for creating a person record. This is a protected endpoint, so you’ll need the `static_token` from the login response and an `X-Action-By` header (typically the admin’s ID).

### Request
Send a `GET` request to the `/api/personinfo/static/static-tables` endpoint with the appropriate headers.

#### Example Request
```http
GET http://localhost:8081/api/personinfo/static/static-tables
Content-Type: application/json
X-Action-By: 083308ac-3c6b-4c5a-a1f5-6dfc63d4d63f
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjA4MzMwOGFjLTNjNmItNGM1YS1hMWY1LTZkZmM2M2Q0ZDYzZiIsImV4cCI6MTc0ODUwMjE3OCwicm9sZSI6MX0.VjRyV3chUPrqOPiC0GK_aLX9nFkY7uiSbyK8-QP0_ew
```

- **URL**: `http://localhost:8081/api/personinfo/static/static-tables`
- **Method**: `GET`
- **Headers**:
  - `Content-Type: application/json`
  - `X-Action-By`: The admin ID (e.g., `"083308ac-3c6b-4c5a-a1f5-6dfc63d4d63f"` from the login response’s `admin.ID`).
  - `Authorization`: The `static_token` from the login response, prefixed with `Bearer` (e.g., `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`).

### Expected Response
A `200 OK` response with the static data.

#### Example Response
```json
HTTP/1.1 200 OK
Access-Control-Allow-Headers: Content-Type, Authorization, Content-Length, Accept-Encoding, Cache-Control, X-Requested-With, accept, origin
Access-Control-Allow-Methods: GET, POST, PUT, DELETE
Access-Control-Allow-Origin: *
Content-Security-Policy: default-src 'none'
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 May 2025 07:03:39 GMT
Content-Length: 627
Connection: close

{
  "blood_groups": [
    {"ID": 1, "Name": "A+"},
    {"ID": 2, "Name": "A-"},
    {"ID": 3, "Name": "B+"},
    {"ID": 4, "Name": "B-"},
    {"ID": 5, "Name": "AB+"},
    {"ID": 6, "Name": "AB-"},
    {"ID": 7, "Name": "O+"},
    {"ID": 8, "Name": "O-"}
  ],
  "religions": [
    {"ID": 1, "ReligionName": "Islam", "ReligionType": "Monotheistic"},
    {"ID": 2, "ReligionName": "Christianity", "ReligionType": "Monotheistic"},
    {"ID": 3, "ReligionName": "Judaism", "ReligionType": "Monotheistic"}
  ],
  "person_types": [
    {"ID": 1, "Type": "Soldier"},
    {"ID": 2, "Type": "Officer"},
    {"ID": 3, "Type": "Civilian"}
  ],
  "ranks": [
    {"ID": 1, "Name": "Private", "DeletedAt": 0},
    {"ID": 2, "Name": "Sergeant", "DeletedAt": 0},
    {"ID": 3, "Name": "Lieutenant", "DeletedAt": 0}
  ]
}
```

- **Status**: `200 OK`
- **Body**:
  - `blood_groups`: List of blood group IDs and names (e.g., `A+`, `O-`).
  - `religions`: List of religion IDs, names, and types (e.g., `Islam`).
  - `person_types`: List of person type IDs and types (e.g., `Soldier`).
  - `ranks`: List of rank IDs and names (e.g., `Private`).

### Action
Use the IDs from this response (e.g., `blood_group_id`, `religion_id`, `person_type_id`, `rank_id`) when creating a new person record in the next step.

---

## Step 3: Create a Full Person Record

The `/api/personinfo/dynamic/persons/full` endpoint allows you to create a new person record with all related data (family info, contact info, skills, etc.) in a single request. This is a protected endpoint, so you’ll need the `dynamic_token` from the login response and an `X-Action-By` header.

### Request
Send a `POST` request to the `/api/personinfo/dynamic/persons/full` endpoint with the person data.

#### Example Request
```http
POST http://localhost:8081/api/personinfo/dynamic/persons/full
Content-Type: application/json
X-Action-By: a93506dc-e9b0-466b-b8ed-0a55b56eee97
Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbl9pZCI6IjA4MzMwOGFjLTNjNmItNGM1YS1hMWY1LTZkZmM2M2Q0ZDYzZiIsImV4cCI6MTc0ODQxNjI5Mywicm9sZSI6MX0.oqItLrWbW7XUjVQict_J-oWpJe17OnT5AT-TGJ9hGBk

{
  "national_id_number": "012345679",
  "first_name": "John",
  "last_name": "Doe",
  "birth_date": "1990-01-01",
  "family_info": {
    "father_details": "{\"name\": \"John Sr.\", \"occupation\": \"Engineer\"}",
    "mother_details": "{\"name\": \"Mary\", \"occupation\": \"Teacher\"}",
    "childs_details": "[]",
    "husband_details": "{}"
  },
  "contact_info": {
    "address": "123 Main St, City, Country",
    "phone_number": "+1234567890",
    "emergency_phone_number": "+0987654321",
    "landline_phone": "123-456-7890",
    "email_address": "john.doe@example.com",
    "social_media": "{\"twitter\": \"@johndoe\", \"linkedin\": \"johndoe\"}"
  },
  "skills": {
    "education": {
      "education_level_id": 1,
      "field_of_study": 1,
      "description": "BSc in Computer Science",
      "university": "Example University",
      "start_date": 1514764800,
      "end_date": 1625097600
    },
    "languages": "[\"English\", \"Spanish\"]",
    "skills_description": "Proficient in Go, Python, and SQL",
    "certificates": "AWS Certified Developer, PMP"
  },
  "physical_info": {
    "blood_group_id": 1,
    "height": 175,
    "weight": 70
  },
  "religion": {
    "religion_id": 1
  },
  "person_type": {
    "person_type_id": 1
  },
  "military_details": {
    "rank_id": 1,
    "service_start_date": 1514764800,
    "service_dispatch_date": 1625097600,
    "service_unit": 101,
    "battalion_unit": 202,
    "company_unit": 303
  }
}
```

- **URL**: `http://localhost:8081/api/personinfo/dynamic/persons/full`
- **Method**: `POST`
- **Headers**:
  - `Content-Type: application/json`
  - `X-Action-By`: A unique identifier for the action (e.g., `"a93506dc-e9b0-466b-b8ed-0a55b56eee97"`).
  - `Authorization`: The `dynamic_token` from the login response, prefixed with `Bearer`.
- **Body**:
  - `national_id_number`: A unique identifier for the person (e.g., `"012345679"`).
  - `first_name` and `last_name`: The person’s name (e.g., `"John"` and `"Doe"`).
  - `birth_date`: The person’s birth date in `YYYY-MM-DD` format (e.g., `"1990-01-01"`).
  - `family_info`: JSON strings for family details (e.g., `father_details`, `mother_details`).
  - `contact_info`: Contact details (e.g., `address`, `phone_number`, `email_address`).
  - `skills`: Skills and education details (e.g., `education`, `languages`).
  - `physical_info`: Physical attributes (e.g., `blood_group_id`, `height`, `weight`).
  - `religion`: The religion ID (e.g., `1` for `Islam`, from the static tables).
  - `person_type`: The person type ID (e.g., `1` for `Soldier`, from the static tables).
  - `military_details`: Military details (e.g., `rank_id`, `service_start_date`).

### Expected Response
A `201 Created` response with the created person’s `national_id_number`.

#### Example Response
```json
HTTP/1.1 201 Created
Access-Control-Allow-Headers: Content-Type, Authorization, Content-Length, Accept-Encoding, Cache-Control, X-Requested-With, accept, origin
Access-Control-Allow-Methods: GET, POST, PUT, DELETE
Access-Control-Allow-Origin: *
Content-Security-Policy: default-src 'none'
Content-Type: application/json; charset=utf-8
X-Content-Type-Options: nosniff
Date: Wed, 28 May 2025 07:05:10 GMT
Content-Length: 34
Connection: close

{
  "national_id_number": "012345689"
}
```

- **Status**: `201 Created`
- **Body**:
  - `national_id_number`: The identifier of the newly created person (e.g., `"012345689"`).

### Notes
- Ensure the IDs (e.g., `blood_group_id`, `religion_id`, `person_type_id`, `rank_id`) match the values from the static tables response.
- The `national_id_number` in the response (`"012345689"`) differs from the request (`"012345679"`), which might indicate a server-side modification or error. Verify this behavior with the API documentation or server logs.

---

## Troubleshooting

### Common Issues and Solutions
1. **401 Unauthorized on Protected Endpoints**:
   - **Cause**: Invalid or expired token in the `Authorization` header.
   - **Solution**: Ensure you’re using the correct token (`static_token` for `GET /static/static-tables`, `dynamic_token` for `POST /dynamic/persons/full`). Re-authenticate if the token has expired.

2. **400 Bad Request on Create Full Person**:
   - **Cause**: Missing or invalid fields in the request body.
   - **Solution**: Check the error message in the response (e.g., `"Key: 'FullPersonRequest.PersonType.PersonTypeID' Error:Field validation for 'PersonTypeID' failed on the 'required' tag"`). Ensure all required fields are present and match the expected JSON keys (e.g., `"person_type_id"` instead of `"type_id"`).

3. **500 Internal Server Error**:
   - **Cause**: Server-side issue, such as a `nil pointer dereference`.
   - **Solution**: Check server logs for details (e.g., `[GIN] 2025/05/28 - 10:14:50 | 500 | ...`). Common causes include uninitialized services or a `nil` database connection. Ensure all services are properly initialized in `main.go`.

4. **Database Connection Issues**:
   - **Cause**: The database (e.g., `test.db`) is inaccessible or not initialized.
   - **Solution**: Verify the database file exists and is accessible. Ensure the database connection is properly set up in your application (e.g., `db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})`).

### Debugging Tips
- **Enable Server Logging**: Add logging in your Go application to trace request handling (e.g., `log.Printf` statements in handlers).
- **Validate Tokens**: Use a JWT decoder (e.g., jwt.io) to verify the `exp` (expiration) claim in your tokens.
- **Check Static Data**: Ensure the static tables contain the expected data by querying the database directly (e.g., `SELECT * FROM blood_groups`).

---

## Conclusion

You’ve successfully interacted with the PersonInfo system API by:
1. Authenticating via the `/api/personinfo/login` endpoint to obtain JWT tokens.
2. Retrieving static data from `/api/personinfo/static/static-tables` using the `static_token`.
3. Creating a new person record via `/api/personinfo/dynamic/persons/full` using the `dynamic_token`.

This guide provides a foundation for using the system. For advanced usage (e.g., updating or deleting records), refer to the API documentation or extend this guide with additional endpoints.

