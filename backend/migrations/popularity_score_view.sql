CREATE MATERIALIZED VIEW POPULARITY_SCORE AS (
       SELECT u.id, u.user_name, 
       -- COUNT(DISTINCT p.id) AS posts_count, 
       -- COUNT(DISTINCT f.post_id) / COALESCE(NULLIF(COUNT(DISTINCT p.id), 0), 1) AS favorites_avg, 
       -- COUNT(DISTINCT c.id) / COALESCE(NULLIF(COUNT(DISTINCT p.id), 0), 1) AS comments_avg,
       (
              COUNT(DISTINCT f.post_id) / COALESCE(NULLIF(COUNT(DISTINCT p.id), 0), 1)     * 2
              + COUNT(DISTINCT c.id) / COALESCE(NULLIF(COUNT(DISTINCT p.id), 0), 1)        * 2
              + COUNT(DISTINCT f.post_id)                                                  * 1.5
              + COUNT(DISTINCT c.id)                                                       * 1.5
       ) AS score
       FROM users u
       LEFT JOIN posts p ON p.author_id = u.id
       LEFT JOIN favorites f ON f.post_id = p.id AND f.user_id != u.id
       LEFT JOIN comments c ON c.post_id = p.id AND c.author_id != u.id
       GROUP BY u.id
       ORDER BY score DESC
);
