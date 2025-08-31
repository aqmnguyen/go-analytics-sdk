export interface Event {
  user_id: string;
  event_type: 'pageview' | 'click';
  event_url: string;
  eventData?: Record<string, any>;
}
