CREATE TABLE analyst_updates (
    update_id       BIGINT     PRIMARY KEY DEFAULT unique_rowid(),
    company_id      BIGINT     NOT NULL,
    brokerage_id    BIGINT     NOT NULL,
    action_id       BIGINT     NOT NULL,
    rating_from_id  BIGINT     NOT NULL,
    rating_to_id    BIGINT     NOT NULL,
    target_from     DECIMAL(12,2) NOT NULL,
    target_to       DECIMAL(12,2) NOT NULL,
    event_time      TIMESTAMP  NOT NULL,

    FOREIGN KEY (company_id)     REFERENCES companies(company_id),
    FOREIGN KEY (brokerage_id)   REFERENCES brokerages(brokerage_id),
    FOREIGN KEY (action_id)      REFERENCES action_types(action_id),
    FOREIGN KEY (rating_from_id) REFERENCES rating_types(rating_id),
    FOREIGN KEY (rating_to_id)   REFERENCES rating_types(rating_id)
);