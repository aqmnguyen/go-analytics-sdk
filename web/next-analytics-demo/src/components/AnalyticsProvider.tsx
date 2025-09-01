'use client';

import { Analytics, trackEvent } from '@/utils/analytics';
import { useEffect } from 'react';
export default function AnalyticsProvider({
  children,
}: {
  children?: React.ReactNode;
}) {
  useEffect(() => {
    // Initialize the analytics instance
    const analytics = Analytics();
    console.log(analytics);
    trackEvent('user_007', 'pageview', {});
  }, []);

  return children ? <div>{children}</div> : null;
}
