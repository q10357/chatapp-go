CREATE TABLE IF NOT EXISTS `user_rel` (
    id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
    user_id_requester BIGINT UNSIGNED NOT NULL,
    user_id_requested BIGINT UNSIGNED NOT NULL,
    status VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    UNIQUE(user_id_requester, user_id_requested), -- Assuming a combination of requester and requested should be unique
    FOREIGN KEY (user_id_requester) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id_requested) REFERENCES user(id) ON DELETE CASCADE ON UPDATE CASCADE
);