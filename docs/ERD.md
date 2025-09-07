# HSI SANDBOX LMS DATABASE ERD

## Entity Relationship Diagram

[CLICK TO VIEW IN MERMAID LIVE EDITOR](https://mermaid.live/edit#pako:eNrFV-FP4jAU_1eaJgZNhnEoOPeNA4yLCIYNc7mQkLpVbGTt0nV3Ivq_X7uN0d0Gp3cfCCNrX3_v9b1fX1-7NfRZgKENm83mjAoisA1m8MZ1gIto8MhewfDOncEZTcdnFPM-QQuOwhkFYOoOJi54f2822TrtzCfj4cAFNkBxTBY0VqBMVAci1F8mAY4rprxJt3erED7HSOCgYkZ15veDyZ3jus545NrgGaVmNJnEAgXeic1nyWGgd9O995QTBaCQbCDSA1e3UAF4Xfd2O7yJKRtzp9-qLij8XkTZRury3Bm53mTa88aT2ljGO3HqtzG5Vk0A_IQEQD73t1k_FpzQBaAoxCUBDhFZZhKBXwVAQcBxHJcwC0wDzGX6DEbTu-NGiJbYAE9YvRsnKodSbRLiWKAwegNJFKjVnSMhdUZjD4ymw6EB-oPr7nToAcp-Hdep5UnxCbWPLOAsdb4YsMxLn5NIEEYP5ri2W3TvkxjzuXxf32pCzpZYF2rzZpsxnbgUYiF_XGXyoyPgszBisSwDkiBwnM9kbKyfaM7pW-3_uNWWaf4Jq7XhpsII85DIoBjdwcR24fYszg4uSsZrGck34X4yfPZTcRqiRZkTWXqXWNtglfyrXfoDJGVR9PYtj-DIf9EdrUZJqACMy4JxsEjycr43EP8ZRaLM-SaUVYSLWted9p0xeAcPTn-g3t7gu7eteTuW2GdUYJrnG03CR1k8D8tIdh79Ix9_S-BUSqiEJ76Sxts8QELgMBLzJQlJuUa9EBpoztcEGSR4rtg5GGnasb1_T6D4pVKyKhtakZTwZZUbykrESG9FEhcJGHD0JBpGI04eJYXSf9mW17RANk7An-wps7HPeJWzVFpbH79QOnV7G3fyQb1SaveTfbTVnnab-rKPkm7Pcx4GDQM0nFHePswtBBpwwUkAbZn52IChPEiQ6sI07BkUz1gekFBdvAPEX5SW0okQ_cFYuFHjLFk8Q_sJLWPZy7zO7-KFFCWCuSvqFzrpjazHEiqgbbXNdmoV2mv4Cu1Wu3Xa6VgXpnl1dnVpnnUuDLiCtmm2T80z6_L8_MLqtC8s6_zDgG-pI-apabUt-W-3ri7PLKvTMiAOiGD8LvuISL8lPn4DwJrp6g)

this is the minimal ERD for the HSI Sandbox LMS, it's currently incomplete and is work in progress.

```mermaidjs
---
tite: "HSI Sandbox LMS"
---

erDiagram
  USERS ||--o{ USER_ROLES : assigns
  ROLES ||--o{ USER_ROLES : includes
  USERS ||--o{ TRACKS : created
  ROLES ||--o{ ROLE_PERMISSIONS: has
  PERMISSIONS || --o{ROLE_PERMISSIONS: has
  TRACKS || --o{ CHAPTERS : has
  CHAPTERS || --o{ LESSONS: has
  CHAPTERS || --o{ TASKS: has
  USERS || --o{ SUBMISSIONS: has
  TASKS || --o{ SUBMISSIONS: has
  USERS || --o{ TRACK_INSTRUCTOR: has
  TRACKS || --O{ TRACK_INSTRUCTOR: has


  USERS {
    cuid id PK
    string name
    string email
    text address
    string gender "ENUM('male, female')"
    timestampz updated_at "NOT NULL, DEFAULT now()"
    timestampz created_at "NOT NULL, DEFAULT now()"
  }

  ROLES {
    cuid id PK
    string name
    string description
    timestampz updated_at "NOT NULL, DEFAULT now()"
    timestampz created_at "NOT NULL, DEFAULT now()"
  }

  USER_ROLES {
    cuid user_id FK
    cuid role_id FK
    timestampz assigned_at
    string assigned_by
    %% composite PK (user_id, role_id)
  }

  PERMISSIONS {
    cuid id PK
    string name
    string description
  }

  ROLE_PERMISSIONS {
    cuid id PK
    cuid role_id FK
    cuid permission_id FK
    timestampz updated_at
    timestampz created_at
    %% composite PK (permission_id, role_id)
  }

  TRACKS {
    cuid id PK
    string cover_image
    string title
    text description
    cuid user_id FK
    timestampz updated_at "NOT NULL, DEFAULT now()"
    timestampz created_at "NOT NULL, DEFAULT now()"
  }

  CHAPTERS {
    cuid id PK
    cuid track_id FK
    string title
    int order
    timestampz updated_at "NOT NULL, DEFAULT now()"
    timestampz created_at "NOT NULL, DEFAULT now()"
  }

  LESSONS {
    cuid id PK
    cuid chapter_id FK
    string type "ENUM('AUDIO | VIDEO | TEXT')"
    string title
    text content
    number order
    timestampz updated_at "NOT NULL, DEFAULT now()"
    timestampz created_at "NOT NULL, DEFAULT now()"
  }

  TASKS {
    cuid id PK
    cuid chapter_id FK
    string title
    text description
    text instructions
    int attempt_limit
    string kind "NOT NULL"
    timestampz due_date
    timestampz updated_at "NOT NULL, DEFAULT now()"
    timestampz created_at "NOT NULL, DEFAULT now()"
  }

  SUBMISSIONS {
    cuid id PK
    cuid task_id FK
    cuid user_id FK
    text url
    int attempt_no
    string status "ENUM('draft','submitted','graded') NOT NULL"
    int score
    timestampz scored_at
    timestampz updated_at
    timestampz created_at
    timestampz submitted_at
  }

  TRACK_INSTRUCTOR{
    cuid id PK
    cuid user_id FK
    cuid track_id
    string status "ENUM('ACTIVE', 'INACTIVE')"
    timestampz updated_at "NOT NULL, DEFAULT now()"
    timestampz created_at "NOT NULL, DEFAULT now()"
  }
```
