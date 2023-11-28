CREATE TABLE image
(
    id         INT NOT NULL AUTO_INCREMENT,
    ad_id      INT NOT NULL,
    content    BINARY,
    'order'    INT NOT NULL,
    created_at TIMESTAMP
);