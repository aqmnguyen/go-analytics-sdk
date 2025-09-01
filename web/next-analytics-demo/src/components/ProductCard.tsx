'use client';

import Image from 'next/image';
import { trackEvent } from '@/utils/analytics';

interface ProductCardProps {
  id: string;
  name: string;
  price: string;
  description: string;
  image: string;
}

export default function ProductCard({
  id,
  name,
  price,
  description,
  image,
}: ProductCardProps) {
  const handleAddToCart = () => {
    trackEvent('user_007', 'click', {
      element: 'add_to_cart_button',
      product_id: id,
      product_name: name,
      product_price: price,
    });
  };

  const handleViewDetails = () => {
    trackEvent('user_007', 'click', {
      element: 'view_details_button',
      product_id: id,
      product_name: name,
    });
  };

  return (
    <div className='bg-white rounded-lg shadow-md p-6 max-w-sm mx-auto h-full flex flex-col'>
      <div className='w-full h-48 relative rounded-lg mb-4 overflow-hidden'>
        <Image
          src={image}
          alt={name}
          fill
          className='object-cover'
          sizes='(max-width: 768px) 100vw, (max-width: 1200px) 50vw, 33vw'
        />
      </div>

      <h3 className='text-xl font-semibold mb-2 text-gray-900'>{name}</h3>
      <p className='text-gray-600 mb-4 flex-grow'>{description}</p>
      <p className='text-2xl font-bold text-green-600 mb-4'>{price}</p>

      <div className='space-y-2 mt-auto'>
        <button
          onClick={handleAddToCart}
          className='w-full bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-700 hover:shadow-lg hover:scale-105 transition-all duration-200 cursor-pointer'
        >
          Add to Cart
        </button>
        <button
          onClick={handleViewDetails}
          className='w-full bg-gray-200 text-gray-800 py-2 px-4 rounded-lg hover:bg-gray-300 hover:shadow-md hover:scale-105 transition-all duration-200 cursor-pointer'
        >
          View Details
        </button>
      </div>
    </div>
  );
}
