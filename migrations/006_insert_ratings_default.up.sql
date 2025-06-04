-- default values for ratings

INSERT INTO rating_types (code, description) VALUES
    ('Outperform',      'Expected to perform better than the market'),
    ('Buy',             'Recommendation to purchase the asset'),
    ('Neutral',         'Expected to perform in line with the market'),
    ('Market Perform',  'Expected to perform in line with the broader market')
;