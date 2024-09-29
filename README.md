# life-journal
ER Diagram


# Business Context

## User Table

**Description**: Table for multi-user support.

| Column Name     | Data Type        | Constraints                        | Description                                                        |
|-----------------|------------------|------------------------------------|--------------------------------------------------------------------|
| `id`            | `UUID`           | `NOT NULL`, `DEFAULT uuidv7`       | Unique identifier for each user, generated automatically.           |
| `created_at`    | `TIMESTAMP`      | `DEFAULT CURRENT_TIMESTAMP`        | Timestamp when the record was created.                             |
| `updated_at`    | `TIMESTAMP`      | `DEFAULT CURRENT_TIMESTAMP`        | Timestamp when the record was last updated.                        |
| `deleted_at`    | `TIMESTAMP`      | `DEFAULT NULL`                     | Timestamp for soft deletes, marking the record as deleted without actual removal. |
| `username`      | `VARCHAR(100)`   | `UNIQUE`                           | Username for the user.                                             |
| `email`         | `VARCHAR(100)`   | `UNIQUE`                           | User's email address.                                              |
| `password`      | `VARCHAR(255)`   |                                    | Hashed password.                                                   |

**Notes**:
- `id`: Unique identifier for each user, generated automatically.
- `created_at`: Timestamp when the record was created.
- `updated_at`: Timestamp when the record was last updated.
- `deleted_at`: Soft delete mechanism to mark the record as deleted without actual removal.
- `username`: Username for the user.
- `email`: User's email address.
- `password`: Hashed password.

```sql
Table User [NOTE: "table for multi-user support."] {
  id UUID [not null, default: `uuidv7`, NOTE: "Unique identifier for each user, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  username VARCHAR(100) [unique, NOTE: "Username for the user."]
  email VARCHAR(100) [unique, NOTE: "User's email address."]
  password VARCHAR(255) [NOTE: "Hashed password."]
}
```

## UserJournal Table

**Description**: Table to link journals to users.

| Column Name   | Data Type      | Constraints                        | Description                                                        |
|---------------|----------------|------------------------------------|--------------------------------------------------------------------|
| `id`          | `UUID`         | `NOT NULL`, `DEFAULT uuidv7`       | Unique identifier for each entry, generated automatically.          |
| `created_at`  | `TIMESTAMP`    | `DEFAULT CURRENT_TIMESTAMP`        | Timestamp when the record was created.                             |
| `updated_at`  | `TIMESTAMP`    | `DEFAULT CURRENT_TIMESTAMP`        | Timestamp when the record was last updated.                        |
| `deleted_at`  | `TIMESTAMP`    | `DEFAULT NULL`                     | Timestamp for soft deletes, marking the record as deleted without actual removal. |
| `user_id`     | `UUID`         | `REFERENCES User(id)`              | References the user who owns the journal.                          |
| `journal_id`  | `UUID`         | `REFERENCES Journal(id)`           | References the journal entry.                                      |

**Notes**:
- `id`: Unique identifier for each entry, generated automatically.
- `created_at`: Timestamp when the record was created.
- `updated_at`: Timestamp when the record was last updated.
- `deleted_at`: Soft delete mechanism to mark the record as deleted without actual removal.
- `user_id`: Foreign key referencing the user who owns the journal.
- `journal_id`: Foreign key referencing the journal entry.

```sql
Table UserJournal [NOTE: "UserJournal to link journals to users."] {
  id UUID [not null, default: `uuidv7`, NOTE: "Unique identifier for each entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  user_id UUID [ref: > User.id, NOTE: "References the user who owns the journal."]
  journal_id UUID [ref: > Journal.id, NOTE: "References the journal entry."]
}
```


## Mood Table

**Description**: Table for storing mood details.

| Column Name | Data Type  | Constraints                | Description                  |
|-------------|------------|----------------------------|------------------------------|
| `id`        | `UUID`     | `PRIMARY KEY`, `UNIQUE`     | Unique identifier for each mood entry. |
| `name`      | `VARCHAR`  |                            | Name of the mood.            |

**Notes**:
- `id`: Unique identifier for each mood entry.
- `name`: Name representing the mood.

```sql
Table Mood {
  id UUID [pk, unique, NOTE: "Unique identifier for each mood entry."]
  name VARCHAR [NOTE: "Name of the mood."]
}
```


## JournalEntry Table

**Description**: JournalEntry to store the detailed content of journal entries.

| Column Name   | Data Type   | Constraints                        | Description                                                        |
|---------------|-------------|------------------------------------|--------------------------------------------------------------------|
| `id`          | `UUID`      | `NOT NULL`, `DEFAULT uuidv7`       | Unique identifier for each journal entry, generated automatically.  |
| `created_at`  | `TIMESTAMP` | `DEFAULT CURRENT_TIMESTAMP`        | Timestamp when the record was created.                             |
| `updated_at`  | `TIMESTAMP` | `DEFAULT CURRENT_TIMESTAMP`        | Timestamp when the record was last updated.                        |
| `deleted_at`  | `TIMESTAMP` | `DEFAULT NULL`                     | Timestamp for soft deletes, marking the record as deleted without actual removal. |
| `journal_id`  | `UUID`      | `REFERENCES Journal(id)`           | References the journal this entry belongs to.                      |
| `topic_id`    | `UUID`      | `REFERENCES Topics(id)`            | References the topic this entry belongs to.                        |
| `entry_text`  | `TEXT`      |                                    | The content of the journal entry.                                  |

**Notes**:
- `id`: Unique identifier for each journal entry, generated automatically.
- `created_at`: Timestamp when the record was created.
- `updated_at`: Timestamp when the record was last updated.
- `deleted_at`: Soft delete mechanism to mark the record as deleted without actual removal.
- `journal_id`: Foreign key referencing the journal this entry belongs to.
- `topic_id`: Foreign key referencing the topic this entry belongs to.
- `entry_text`: Content of the journal entry.

```sql
Table JournalEntry [NOTE: "JournalEntry to store the detailed content of journal entries."] {
  id UUID [not null, default: `uuidv7`, NOTE: "Unique identifier for each journal entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  journal_id UUID [ref: > Journal.id, NOTE: "References the journal this entry belongs to."]
  topic_id UUID [ref: > Topics.id, NOTE: "References the topic this entry belongs to."]
  entry_text TEXT [NOTE: "The content of the journal entry."]
}
```

## Topics Table

**Description**: A table to store different journal topics.

| Column Name | Data Type     | Constraints      | Description                   |
|-------------|---------------|-----------------|-------------------------------|
| `id`        | `UUID`        | `PRIMARY KEY`, `UNIQUE` | Unique identifier for each topic. |
| `name`      | `VARCHAR(255)`|                 | The name of the topic.         |

**Notes**:
- `id`: Unique identifier for each topic.
- `name`: The name of the topic.

```sql
Table Topics [NOTE: "A table to store different journal topics."] {
  id UUID [pk, unique, NOTE: "Unique identifier for each topic."]
  name VARCHAR(255) [NOTE: "The name of the topic."]
}
```

## Tag Table

**Description**: Table for storing tags used for topic-based tagging of journals.

| Column Name  | Data Type     | Constraints                         | Description                                              |
|--------------|---------------|-------------------------------------|----------------------------------------------------------|
| `id`         | `UUID`        | `NOT NULL`, `DEFAULT uuidv7`, `PRIMARY KEY` | Unique identifier for each tag, generated automatically.   |
| `created_at` | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was created.                    |
| `updated_at` | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was last updated.               |
| `deleted_at` | `TIMESTAMP`   | `DEFAULT NULL`                      | Timestamp for soft deletes, marking the record as deleted.|
| `name`       | `VARCHAR(100)`| `UNIQUE`                            | Name of the tag.                                          |

**Notes**:
- `id`: Unique identifier for each tag, generated automatically.
- `name`: The name of the tag, which must be unique.

```sql
Table Tag [NOTE: "Tags & JournalTag for topic-based tagging of journals."] {
  id UUID [not null, default: `uuidv7`, pk, unique, NOTE: "Unique identifier for each tag, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted."]
  name VARCHAR(100) [unique, NOTE: "Name of the tag."]
}
```

## JournalTag Table

**Description**: Table for linking tags to journal entries for topic-based tagging of journals.

| Column Name  | Data Type     | Constraints                         | Description                                              |
|--------------|---------------|-------------------------------------|----------------------------------------------------------|
| `id`         | `UUID`        | `NOT NULL`, `DEFAULT uuidv7`, `PRIMARY KEY` | Unique identifier for each tag entry, generated automatically.   |
| `created_at` | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was created.                    |
| `updated_at` | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was last updated.               |
| `deleted_at` | `TIMESTAMP`   | `DEFAULT NULL`                      | Timestamp for soft deletes, marking the record as deleted.|
| `journal_id` | `UUID`        | `REFERENCES Journal(id)`            | References the journal.                                   |
| `tag_id`     | `UUID`        | `REFERENCES Tag(id)`                | References the tag.                                       |

**Notes**:
- `id`: Unique identifier for each journal tag entry, generated automatically.
- `journal_id`: References the specific journal entry.
- `tag_id`: References the specific tag applied to the journal.

```sql
Table JournalTag [NOTE: "Tags & JournalTag for topic-based tagging of journals."] {
  id UUID [not null, default: `uuidv7`, pk, unique, NOTE: "Unique identifier for each tag entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted."]
  journal_id UUID [ref: > Journal.id, NOTE: "References the journal."]
  tag_id UUID [ref: > Tag.id, NOTE: "References the tag."]
}
```


## Attachment Table

**Description**: Table for uploading files to journals.

| Column Name  | Data Type     | Constraints                         | Description                                              |
|--------------|---------------|-------------------------------------|----------------------------------------------------------|
| `id`         | `UUID`        | `NOT NULL`, `DEFAULT uuidv7`, `PRIMARY KEY` | Unique identifier for each upload entry, generated automatically.   |
| `created_at` | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was created.                    |
| `updated_at` | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was last updated.               |
| `deleted_at` | `TIMESTAMP`   | `DEFAULT NULL`                      | Timestamp for soft deletes, marking the record as deleted.|
| `journal_id` | `UUID`        | `REFERENCES Journal(id)`            | References the journal to which the attachment belongs.   |
| `file_path`  | `VARCHAR(255)`|                                      | File path or URL of the attachment.                       |
| `file_type`  | `VARCHAR(50)` |                                      | The type of the file (e.g., image, document).            |

**Notes**:
- `id`: Unique identifier for each attachment entry, generated automatically.
- `journal_id`: References the specific journal entry associated with the attachment.
- `file_path`: Specifies the location of the uploaded file.
- `file_type`: Indicates the file format.

```sql
Table Attachment [NOTE: "Attachment for uploading files to journals."] {
  id UUID [not null, default: `uuidv7`, pk, unique, NOTE: "Unique identifier for each upload entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted."]
  journal_id UUID [ref: > Journal.id, NOTE: "References the journal to which the attachment belongs."]
  file_path VARCHAR(255) [NOTE: "File path or URL of the attachment."]
  file_type VARCHAR(50) [NOTE: "The type of the file (e.g., image, document)."]
}
```

## Reminder Table

**Description**: Reminder to set reminders for journal entries.

| Column Name      | Data Type     | Constraints                         | Description                                           |
|------------------|---------------|-------------------------------------|-------------------------------------------------------|
| `id`             | `UUID`        | `NOT NULL`, `DEFAULT uuidv7`, `PRIMARY KEY` | Unique identifier for each reminder entry, generated automatically. |
| `created_at`     | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was created.                 |
| `updated_at`     | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was last updated.            |
| `deleted_at`     | `TIMESTAMP`   | `DEFAULT NULL`                      | Timestamp for soft deletes, marking the record as deleted. |
| `journal_id`     | `UUID`        | `REFERENCES Journal(id)`            | References the journal entry associated with the reminder. |
| `reminder_date`  | `TIMESTAMP`   |                                     | Date and time when the reminder will trigger.          |
| `message`        | `TEXT`        |                                     | Optional reminder message.                             |

**Notes**:
- `id`: Unique identifier for each reminder entry, generated automatically.
- `journal_id`: References the specific journal entry associated with the reminder.
- `reminder_date`: Specifies when the reminder should trigger.
- `message`: Allows for an optional message to be included with the reminder.

```sql
Table Reminder [NOTE: "Reminder to set reminders for journal entries."] {
  id UUID [not null, default: `uuidv7`, pk, unique, NOTE: "Unique identifier for each reminder entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted."]
  journal_id UUID [ref: > Journal.id, NOTE: "References the journal entry."]
  reminder_date TIMESTAMP [NOTE: "Date and time when the reminder will trigger."]
  message TEXT [NOTE: "Optional reminder message."]
}
```

## Journal Table

**Description**: Journal to store journal entries and related information.

| Column Name      | Data Type     | Constraints                         | Description                                           |
|------------------|---------------|-------------------------------------|-------------------------------------------------------|
| `id`             | `UUID`        | `NOT NULL`, `DEFAULT uuidv7`, `PRIMARY KEY` | Unique identifier for each journal entry, generated automatically. |
| `created_at`     | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was created.                 |
| `updated_at`     | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was last updated.            |
| `deleted_at`     | `TIMESTAMP`   | `DEFAULT NULL`                      | Timestamp for soft deletes, marking the record as deleted. |
| `name`           | `VARCHAR(100)`|                                     | Name of the journal.                                  |
| `favorite`       | `BOOLEAN`     |                                     | Indicates whether the journal is marked as favorite.  |
| `category_id`    | `UUID`        |                                     | References the category associated with the journal.   |
| `date`           | `DATE`        |                                     | The date associated with the journal entry.           |
| `archive_status` | `BOOLEAN`     | `DEFAULT false`                     | Indicates whether the journal is archived.            |

**Notes**:
- `id`: Unique identifier for each journal entry, generated automatically.
- `name`: Specifies the name of the journal.
- `favorite`: A boolean flag to indicate if the journal is a favorite.
- `category_id`: Can link to a category table for better organization.
- `date`: Stores the date related to the journal entry.
- `archive_status`: Shows if the journal is archived (default is false).

```sql
Table Journal [NOTE: "Journal"] {
  id UUID [not null, default: `uuidv7`, pk, unique, NOTE: "Unique identifier for each journal entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted."]
  name VARCHAR(100) [NOTE: ""]
  favorite BOOLEAN
  category_id UUID 
  date DATE
  archive_status BOOLEAN [default: false, NOTE: "Indicates whether the journal is archived."]
}
```

## MoodLog Table

**Description**: MoodLog to record user moods along with timestamps.

| Column Name      | Data Type     | Constraints                         | Description                                           |
|------------------|---------------|-------------------------------------|-------------------------------------------------------|
| `id`             | `UUID`        | `NOT NULL`, `DEFAULT uuidv7`, `PRIMARY KEY` | Unique identifier for each mood log entry, generated automatically. |
| `created_at`     | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was created.                 |
| `updated_at`     | `TIMESTAMP`   | `DEFAULT CURRENT_TIMESTAMP`         | Timestamp when the record was last updated.            |
| `deleted_at`     | `TIMESTAMP`   | `DEFAULT NULL`                      | Timestamp for soft deletes, marking the record as deleted. |
| `user_id`        | `UUID`        | `REF: > User.id`                   | References the user who logged the mood.               |
| `mood_id`        | `UUID`        | `REF: > Mood.id`                   | References the mood being logged.                       |

**Notes**:
- `id`: Unique identifier for each mood log entry, generated automatically.
- `user_id`: Links the mood log entry to a specific user.
- `mood_id`: Links to a specific mood from the Mood table.

```sql
Table MoodLog {
  id UUID [not null, default: `uuidv7`, pk, unique, NOTE: "Unique identifier for each mood log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, NOTE: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, NOTE: "Timestamp for soft deletes, marking the record as deleted."]
  user_id UUID [ref: > User.id, NOTE: "References the user who logged the mood."]
  mood_id UUID [ref: > Mood.id, NOTE: "References the mood being logged."]
}
```
