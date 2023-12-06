DROP VIEW quiz_session_view;

CREATE VIEW quiz_session_view
AS
SELECT q.sha1                                                                  AS quiz_sha1,
       q.name                                                                  AS quiz_name,
       q.filename                                                              AS quiz_filename,
       q.version                                                               AS quiz_version,
       q.duration                                                              AS quiz_duration,
       q.created_at                                                            AS quiz_created_at,
       CASE WHEN s.uuid IS NULL THEN '' ELSE s.uuid END                        AS session_uuid,
       CASE WHEN s.user_id IS NULL THEN '' ELSE s.user_id END                  AS user_id,
       CASE WHEN u.name IS NULL THEN '' ELSE u.name END                        AS user_name,
       CASE WHEN u.picture IS NULL THEN '' ELSE u.picture END                  AS user_picture,
       CASE WHEN sc.uuid IS NULL THEN '' ELSE sc.uuid END                      AS class_uuid,
       CASE WHEN sc.name IS NULL THEN '' ELSE sc.name END                      AS class_name,
       CASE WHEN sv.remaining_sec IS NULL THEN 0 ELSE sv.remaining_sec END     AS remaining_sec,
       CASE WHEN sv.checked_answers IS NULL THEN 0 ELSE sv.checked_answers END AS checked_answers,
       CASE WHEN sv.results IS NULL THEN 0 ELSE sv.results END                 AS results
FROM quiz q
         LEFT JOIN session s ON q.sha1 = s.quiz_sha1
         LEFT JOIN user u ON s.user_id = u.id
         LEFT JOIN student_class sc ON u.class_uuid = sc.uuid
         LEFT JOIN session_view sv ON q.sha1 = sv.quiz_sha1 AND s.uuid = sv.uuid
WHERE q.active = TRUE;
