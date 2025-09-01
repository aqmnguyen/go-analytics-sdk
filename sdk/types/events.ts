export interface Event {
  user_id: string;
  event_type: 'pageview' | 'click';
  event_url: string;
  event_data?: Record<string, any>;
}
