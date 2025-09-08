
import React from 'react';

const ProductItem = ({ product, onEdit, onDelete }) => {
  return (
    <div className="bg-white shadow-md rounded-lg p-4 flex flex-col justify-between">
      <div>
        <img src={product.imageUrl} alt={product.name} className="w-full h-32 object-cover rounded-t-lg" />
        <h2 className="text-xl font-bold mt-2">{product.name}</h2>
        <p className="text-gray-600">{product.description}</p>
        <p className="text-lg font-semibold mt-2">${product.price/100}</p>
      </div>
      <div className="flex justify-end mt-4">
        <button onClick={() => onEdit(product)} className="bg-blue-500 text-white px-4 py-2 rounded-md mr-2">Edit</button>
        <button onClick={() => onDelete(product.id)} className="bg-red-500 text-white px-4 py-2 rounded-md">Delete</button>
      </div>
    </div>
  );
};

export default ProductItem;
