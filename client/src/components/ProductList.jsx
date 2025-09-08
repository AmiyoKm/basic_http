
import React from 'react';
import ProductItem from './ProductItem';

const ProductList = ({ products, onEdit, onDelete }) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 p-4">
      {products.map(product => (
        <ProductItem key={product.id} product={product} onEdit={onEdit} onDelete={onDelete} />
      ))}
    </div>
  );
};

export default ProductList;
