'use client';

import { trackEvent } from '@/utils/analytics';
import { useCartStore } from '@/utils/cartStore';

export default function CartIcon() {
  const { items, clearCart } = useCartStore();
  const cartTotal = items.reduce(
    (acc, item) => acc + parseFloat(item.price),
    0
  );


  const handleCheckout = () => {
    const orderId = Math.random().toString(36).substring(2, 8);
    // Track checkout event
    trackEvent('user_007', 'click', {
      element: 'checkout_button',
    });

    trackEvent('user_007', 'conversion', {
      order_id: `${orderId}`,
      order_total: cartTotal.toFixed(2),
      products: items.map((item) => ({
        product_id: item.id,
        product_name: item.name,
        product_price: item.price,
        product_quantity: item.quantity,
        product_category: item.category,
      })),
    });

    // Clear cart after checkout
    clearCart();

    // Show success message (you could replace this with a modal or redirect)
    alert(`Checkout successful! Total: $${cartTotal.toFixed(2)}`);
  };

  return (
    <div className='fixed top-4 right-4 bg-white rounded-lg shadow-lg p-4 z-50'>
      <div className='flex items-center gap-4'>
        {/* Cart Icon with Item Count */}
        <div className='relative'>
          <div className='w-8 h-8 bg-blue-600 rounded-full flex items-center justify-center'>
            <span className='text-white text-sm font-semibold'>
              {items.length}
            </span>
          </div>
          {items.length > 0 && (
            <div className='absolute -top-1 -right-1 w-4 h-4 bg-red-500 rounded-full flex items-center justify-center'>
              <span className='text-white text-xs font-bold'>
                {items.length}
              </span>
            </div>
          )}
        </div>

        {/* Cart Info */}
        <div className='text-sm'>
          <div className='font-semibold text-gray-900'>
            {items.length} {items.length === 1 ? 'item' : 'items'}
          </div>
          <div className='text-gray-600'>Total: ${cartTotal.toFixed(2)}</div>
        </div>

        {/* Checkout Button */}
        <button
          onClick={handleCheckout}
          disabled={items.length === 0}
          className='bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 disabled:bg-gray-300 disabled:cursor-not-allowed transition-colors'
        >
          Checkout
        </button>
      </div>
    </div>
  );
}
