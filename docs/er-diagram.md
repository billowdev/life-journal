

Table User [note: "table for multi-user support."] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  username VARCHAR(100) [unique, NOTE: "Username for the user."]
  email VARCHAR(100) [unique, NOTE: "User's email address."]
  password VARCHAR(255) [NOTE: "Hashed password."]
}

Table UserJournal [note: "UserJournal to link journals to users."] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  user_id UUID [ref: > User.id, NOTE: "References the user who owns the journal."]
  journal_id UUID [ref: > Journal.id, NOTE: "References the journal entry."]
}

Table Mood {
    id UUID [pk, unique, NOTE: ""]
    name varchar
}

Table JournalEntry [note: "JournalEntry to store the detailed content of journal entries."] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  journal_id UUID [ref: > Journal.id, NOTE: "References the journal this entry belongs to."]
  topic_id UUID [ref: > Topics.id, NOTE: "References the topic this entry belongs to."]
  entry_text TEXT [NOTE: "The content of the journal entry."]
}

Table Topics {
  id UUID [pk, unique, NOTE: ""]
  name varchar(255)
}


Table Tag [note: "Tags & JournalTag for topic-based tagging of journals."] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  name VARCHAR(100) [unique, NOTE: "Name of the tag."]
}

Table JournalTag [note: "Tags & JournalTag for topic-based tagging of journals."] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  journal_id UUID [ref: > Journal.id, NOTE: "References the journal."]
  tag_id UUID [ref: > Tag.id, NOTE: "References the tag."]
}

Table Attachment [note: "Attachment for uploading files to journals."] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]


  journal_id UUID [ref: > Journal.id, NOTE: "References the journal to which the attachment belongs."]
  file_path VARCHAR(255) [NOTE: "File path or URL of the attachment."]
  file_type VARCHAR(50) [NOTE: "The type of the file (e.g., image, document)."]
}

Table Reminder [note: "Reminder to set reminders for journal entries."] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  journal_id UUID [ref: > Journal.id, NOTE: "References the journal entry."]
  reminder_date TIMESTAMP [NOTE: "Date and time when the reminder will trigger."]
  message TEXT [NOTE: "Optional reminder message."]
}

Table Journal [NOTE: "Journal"] {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  name VARCHAR(100) [NOTE: ""]
  favorite boolean
  category_id uuid 
  date date
  archive_status BOOLEAN [default: false, NOTE: "Indicates whether the journal is archived."]
}

Table MoodLog {
  id uuid [not null, default: `uuidv7`, note: "Unique identifier for each upload log entry, generated automatically."]
  created_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was created."]
  updated_at TIMESTAMP [default: `CURRENT_TIMESTAMP`, note: "Timestamp when the record was last updated."]
  deleted_at TIMESTAMP [default: NULL, note: "Timestamp for soft deletes, marking the record as deleted without actual removal."]

  user_id UUID [ref: > User.id, NOTE: "References the user who logged the mood."]
  mood_id UUID [ref: > Mood.id, NOTE: "References the mood being logged."]
}
