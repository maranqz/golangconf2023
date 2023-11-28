CREATE TABLE ad
(
    id                 INT           NOT NULL AUTO_INCREMENT,
    user_id            INT           NOT NULL,
    category_id        INT           NOT NULL,
    title              VARCHAR(255)  NOT NULL,
    description        VARCHAR(2048) NOT NULL,
    contact            JSON,
    phone              INT,
    quantity_available INT           NOT NULL,
    quantity_reserved  INT           NOT NULL,
    created_at         TIMESTAMP,
    updated_at         TIMESTAMP
);