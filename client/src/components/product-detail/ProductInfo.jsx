// components/productDetail/ProductInfo.jsx
import React from "react";

const ProductInfo = ({ product, selectedVariant }) => {
  return (
    <>
      <h1 className="text-2xl font-semibold">{product.name}</h1>
      <p className="text-sm text-muted-foreground">
        {product.category?.name} / {product.subcategory?.name}
      </p>
      <p className="text-gray-600">{product.description}</p>

      <div className="text-xl font-bold text-primary">
        ${selectedVariant.price.toFixed(2)}
      </div>
    </>
  );
};

export default ProductInfo;
