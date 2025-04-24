// pages/productDetail.jsx
import React, { useEffect, useState } from "react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useParams, useNavigate } from "react-router-dom";
import { useProductDetailQuery } from "@/hooks/useProductsQuery";
import ProductInfo from "@/components/product-detail/ProductInfo";
import ProductGallery from "@/components/product-detail/ProductGallery";
import ProductVariantSelector from "@/components/product-detail/ProductVariantSelector";

const Product = () => {
  const { slug } = useParams();
  const navigate = useNavigate();
  const baseSlug = slug.split("+")[0];
  const variantSlug = slug.split("+")[1];

  const {
    isError,
    refetch,
    isLoading,
    data: product,
  } = useProductDetailQuery(baseSlug);

  const [selectedImage, setSelectedImage] = useState(null);
  const [selectedVariant, setSelectedVariant] = useState(null);

  useEffect(() => {
    if (product?.variants?.length) {
      const variant = variantSlug
        ? product.variants.find((v) => v.sku === variantSlug)
        : product.variants[0];
      setSelectedVariant(variant);
      setSelectedImage(variant?.imageUrl || product.images?.[0]);
    }
  }, [product, variantSlug]);

  const handleVariantClick = (variant) => {
    if (variant.sku !== selectedVariant?.sku) {
      navigate(`/products/${product.slug}+${variant.sku}`);
    }
  };

  if (isLoading) return <FetchLoading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;
  if (!product || !selectedVariant) return null;

  const galleryImages = Array.from(
    new Set([
      ...(product.images || []),
      ...product.variants.map((v) => v.imageUrl).filter(Boolean),
    ])
  );

  return (
    <div className="container mx-auto p-6 grid md:grid-cols-2 gap-8">
      <ProductGallery
        images={galleryImages}
        selectedImage={selectedImage}
        onSelectImage={setSelectedImage}
      />

      <div className="space-y-4">
        <ProductInfo product={product} selectedVariant={selectedVariant} />

        <ProductVariantSelector
          product={product}
          selectedVariant={selectedVariant}
          onSelectVariant={handleVariantClick}
        />
      </div>
    </div>
  );
};

export default Product;
