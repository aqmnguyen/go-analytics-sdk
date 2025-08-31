import type { AnalyticsConfig } from '@/types/config';
import type { Event } from '@/types/events';

export class AnalyticsSDK {
  private config: AnalyticsConfig;

  constructor(config: AnalyticsConfig) {
    if (!config.clientId) {
      throw new Error('Client ID is required');
    }
    this.config = config;
  }

  getConfig() {
    return this.config;
  }

  async sendEvent(event: Event) {
    try {
      event.eventData = event.eventData || {};
      this._validateEvent(event);

      const { baseUrl = 'http://localhost:8080', clientId } = this.config;
      const { event_type } = event;

      const url = `${baseUrl}/event/${event_type}`;
      const headers = {
        'Content-Type': 'application/json',
      };

      const response = await fetch(url, {
        method: 'POST',
        headers,
        body: JSON.stringify({ ...event, client_id: clientId }),
      });

      if (!response.ok) {
        throw new Error(
          `Failed to send event: ${response.status} ${response.statusText}`
        );
      }

      return response.json();
    } catch (error) {
      console.error(error);
      return null;
    }
  }

  private _validateEvent(event: Event) {
    const { clientId } = this.config;

    if (!clientId) {
      throw new Error('Client ID is required');
    }

    const { user_id, event_type, event_url, eventData = {} } = event;

    if (!user_id || !event_type || !event_url || !eventData) {
      throw new Error('User ID, event type and event URL are required');
    }

    if (event_type !== 'pageview' && event_type !== 'click') {
      throw new Error('Invalid event type, must be either pageview or click');
    }

    if (typeof eventData !== 'object') {
      throw new Error('Event data must be an object');
    }
  }
}
