-- Seed sample data for analytics testing
-- This script inserts sample click and pageview events

-- Insert sample pageview events
INSERT INTO public.events (user_id, event_type, event_url, event_data) VALUES
('user_001', 'pageview', 'https://example.com/home', '{"page_title": "Homepage", "referrer": "https://google.com", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36", "screen_resolution": "1920x1080", "time_on_page": 45}'),
('user_002', 'pageview', 'https://example.com/products', '{"page_title": "Products", "referrer": "https://example.com/home", "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36", "screen_resolution": "1366x768", "time_on_page": 120}'),
('user_003', 'pageview', 'https://example.com/about', '{"page_title": "About Us", "referrer": "https://example.com/home", "user_agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15", "screen_resolution": "375x667", "time_on_page": 30}'),
('user_001', 'pageview', 'https://example.com/contact', '{"page_title": "Contact", "referrer": "https://example.com/about", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36", "screen_resolution": "1920x1080", "time_on_page": 60}'),
('user_004', 'pageview', 'https://example.com/blog', '{"page_title": "Blog", "referrer": "https://example.com/home", "user_agent": "Mozilla/5.0 (Linux; Android 11; SM-G991B) AppleWebKit/537.36", "screen_resolution": "360x800", "time_on_page": 180}'),
('user_002', 'pageview', 'https://example.com/cart', '{"page_title": "Shopping Cart", "referrer": "https://example.com/products", "user_agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36", "screen_resolution": "1366x768", "time_on_page": 90}'),
('user_005', 'pageview', 'https://example.com/checkout', '{"page_title": "Checkout", "referrer": "https://example.com/cart", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36", "screen_resolution": "1440x900", "time_on_page": 150}'),
('user_003', 'pageview', 'https://example.com/faq', '{"page_title": "FAQ", "referrer": "https://example.com/about", "user_agent": "Mozilla/5.0 (iPhone; CPU iPhone OS 14_7_1 like Mac OS X) AppleWebKit/605.1.15", "screen_resolution": "375x667", "time_on_page": 75}');

-- Insert sample click events
INSERT INTO public.events (user_id, event_type, event_url, event_data) VALUES
('user_001', 'click', 'https://example.com/home', '{"element_id": "cta-button", "element_text": "Get Started", "element_type": "button", "click_coordinates": {"x": 150, "y": 200}, "page_section": "hero"}'),
('user_002', 'click', 'https://example.com/products', '{"element_id": "product-card-1", "element_text": "Premium Widget", "element_type": "div", "click_coordinates": {"x": 300, "y": 400}, "page_section": "product-grid"}'),
('user_003', 'click', 'https://example.com/about', '{"element_id": "team-member-1", "element_text": "John Doe", "element_type": "img", "click_coordinates": {"x": 200, "y": 150}, "page_section": "team"}'),
('user_001', 'click', 'https://example.com/contact', '{"element_id": "contact-form", "element_text": "Submit", "element_type": "button", "click_coordinates": {"x": 250, "y": 500}, "page_section": "form"}'),
('user_004', 'click', 'https://example.com/blog', '{"element_id": "blog-post-1", "element_text": "Read More", "element_type": "a", "click_coordinates": {"x": 400, "y": 300}, "page_section": "blog-list"}'),
('user_002', 'click', 'https://example.com/cart', '{"element_id": "add-to-cart", "element_text": "Add to Cart", "element_type": "button", "click_coordinates": {"x": 350, "y": 450}, "page_section": "product-detail"}'),
('user_005', 'click', 'https://example.com/checkout', '{"element_id": "payment-button", "element_text": "Pay Now", "element_type": "button", "click_coordinates": {"x": 300, "y": 600}, "page_section": "payment"}'),
('user_003', 'click', 'https://example.com/faq', '{"element_id": "faq-item-1", "element_text": "How do I...", "element_type": "h3", "click_coordinates": {"x": 100, "y": 250}, "page_section": "faq-list"}'),
('user_001', 'click', 'https://example.com/home', '{"element_id": "navigation-menu", "element_text": "Products", "element_type": "nav", "click_coordinates": {"x": 800, "y": 50}, "page_section": "header"}'),
('user_002', 'click', 'https://example.com/products', '{"element_id": "filter-dropdown", "element_text": "Price: Low to High", "element_type": "select", "click_coordinates": {"x": 200, "y": 100}, "page_section": "filters"}');

-- Insert some events with different timestamps for time-based analysis
INSERT INTO public.events (user_id, event_type, event_url, event_data, created_at) VALUES
('user_006', 'pageview', 'https://example.com/home', '{"page_title": "Homepage", "referrer": "direct", "user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36", "screen_resolution": "1920x1080", "time_on_page": 25}', NOW() - INTERVAL '2 hours'),
('user_007', 'click', 'https://example.com/products', '{"element_id": "search-box", "element_text": "Search products...", "element_type": "input", "click_coordinates": {"x": 150, "y": 80}, "page_section": "search"}', NOW() - INTERVAL '1 hour'),
('user_008', 'pageview', 'https://example.com/blog', '{"page_title": "Blog", "referrer": "https://example.com/home", "user_agent": "Mozilla/5.0 (Linux; Android 12; Pixel 6) AppleWebKit/537.36", "screen_resolution": "412x915", "time_on_page": 95}', NOW() - INTERVAL '30 minutes');

-- Verify the data was inserted
SELECT 
    event_type,
    COUNT(*) as event_count,
    COUNT(DISTINCT user_id) as unique_users
FROM public.events 
GROUP BY event_type;
