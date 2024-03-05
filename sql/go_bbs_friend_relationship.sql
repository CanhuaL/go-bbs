CREATE TABLE friend_relationship (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT NOT NULL,
    friend_id BIGINT NOT NULL,
    status ENUM('pending', 'accepted', 'rejected', 'blocked') DEFAULT 'pending',
    created_at TIMESTAMP ,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    INDEX user_friend_index (user_id, friend_id),  -- 索引以便快速查询用户的好友列表
    FOREIGN KEY (user_id) REFERENCES user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (friend_id) REFERENCES user(user_id) ON DELETE CASCADE
);
