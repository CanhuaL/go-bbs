CREATE TABLE sms_messages (
    id INT AUTO_INCREMENT PRIMARY KEY,
    sms_type VARCHAR(50) NOT NULL,
    sms_content TEXT NOT NULL,
    phone VARCHAR(20) NOT NULL,
    send_time DATETIME NOT NULL
);

CREATE INDEX idx_phone ON sms_messages (phone);
