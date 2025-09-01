import { AnalyticsSDK } from '../../../../sdk/src';

declare global {
  interface Window {
    analyticsInstance: AnalyticsSDK;
  }
}

export function Analytics(): AnalyticsSDK {
  if (!window.analyticsInstance) {
    window.analyticsInstance = new AnalyticsSDK({
      clientId: '1234567890', // TODO: replace with actual client ID
      clientKey: '1234567890',
      baseUrl: 'http://localhost:8080',
      debug: true,
    });
  }

  return window.analyticsInstance;
}

export function trackEvent(
  user_id: string,
  event_type: string,
  event_data: Record<string, any>
) {
  const analytics = Analytics();
  console.log(analytics);
  const { event_url, page_data } = _getCurrentPageData();
  event_data = {
    ...event_data,
    ...page_data,
  };

  const data = {
    user_id,
    event_type,
    event_url,
    event_data,
  };

  console.log(data);

  analytics.sendEvent(data);
  if (analytics.getConfig().debug) {
    console.log('Event sent:', data);
  }
}

export function _getCurrentPageData() {
  const event_url = window.location.href;
  const referrer = document.referrer;
  const user_agent = navigator.userAgent;

  return { event_url, page_data: { referrer, user_agent } };
}
