import AnalyticsProvider from '@/components/AnalyticsProvider';
import ProductCard from '@/components/ProductCard';
import CartIcon from '@/components/CartIcon';

export default function Home() {
  const products = [
    {
      id: 'product-1',
      name: 'Wireless Headphones',
      price: '$99.99',
      description:
        'Premium wireless headphones with noise cancellation and 30-hour battery life.',
      image: '/images/headphones.jpg',
      category: ['headphones', 'electronics'],
    },
    {
      id: 'product-2',
      name: 'Smart Watch',
      price: '$199.99',
      description:
        'Advanced smartwatch with health tracking, GPS, and water resistance.',
      image: '/images/smartwatch.jpg',
      category: ['smartwatch', 'electronics'],
    },
    {
      id: 'product-3',
      name: 'Bluetooth Speaker',
      price: '$79.99',
      description:
        'Portable Bluetooth speaker with 360-degree sound and 12-hour battery.',
      image: '/images/speaker.jpg',
      category: ['speaker', 'electronics'],
    },
  ];

  return (
    <>
      <AnalyticsProvider />
      <CartIcon />
      <div className='min-h-screen bg-gray-50 py-8'>
        <div className='max-w-7xl mx-auto px-4 sm:px-6 lg:px-8'>
          <div className='text-center mb-12'>
            <h1 className='text-4xl font-bold text-gray-900 mb-4'>
              Analytics SDK Demo
            </h1>
            <p className='text-xl text-gray-600'>
              Click the buttons below to test event tracking
            </p>
          </div>

          <div className='grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8'>
            {products.map((product) => (
              <ProductCard
                key={product.id}
                id={product.id}
                name={product.name}
                price={product.price}
                description={product.description}
                image={product.image}
                category={product.category}
              />
            ))}
          </div>

          <div className='mt-16 text-center'>
            <p className='text-gray-500 text-sm'>
              Check your browser console and Go server logs to see the tracked
              events
            </p>
          </div>
        </div>
      </div>
    </>
  );
}
