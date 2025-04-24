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

  // Ambil slug dasar dan kombinasi opsi dari URL
  const parts = slug.split("+");
  const baseSlug = parts[0];
  const variantOptionsFromUrl = parts.slice(1).reduce((acc, part) => {
    const [key, value] = part.split("-");
    if (key && value) acc[key] = value;
    return acc;
  }, {});

  const {
    isError,
    refetch,
    isLoading,
    data: product,
  } = useProductDetailQuery(baseSlug);

  const [selectedVariant, setSelectedVariant] = useState(null);
  const [selectedImage, setSelectedImage] = useState(null);
  const [selectedOptions, setSelectedOptions] = useState({});

  useEffect(() => {
    if (product?.variants?.length) {
      // Cari variant berdasarkan opsi di URL atau fallback
      const variant =
        product.variants.find((v) =>
          Object.entries(variantOptionsFromUrl).every(
            ([k, vOpt]) => v.options?.[k] === vOpt
          )
        ) || product.variants[0];

      setSelectedVariant(variant);
      setSelectedOptions({ ...variant.options });
      setSelectedImage(variant?.imageUrl || product.images?.[0]);
    }
  }, [product, slug]);

  // Saat user pilih opsi variant (misal color atau size)
  const handleVariantOptionChange = (key, value) => {
    const updatedOptions = { ...selectedOptions, [key]: value };

    // Jika ganti warna, reset size ke yang tersedia
    if (key === "colors") {
      const availableSizes = product.variants
        .filter((v) => v.options?.colors === value)
        .map((v) => v.options?.["clothing size"]);
      if (availableSizes.length > 0) {
        updatedOptions["clothing size"] = availableSizes[0];
      }
    }

    const matchedVariant = product.variants.find((v) =>
      Object.entries(updatedOptions).every(
        ([k, vOpt]) => v.options?.[k] === vOpt
      )
    );

    if (matchedVariant) {
      setSelectedVariant(matchedVariant);
      setSelectedOptions({ ...matchedVariant.options });
      setSelectedImage(matchedVariant?.imageUrl || product.images?.[0]);

      const variantSlug = Object.entries(matchedVariant.options || {})
        .map(([k, v]) => `${k}-${v}`)
        .join("+");
      navigate(`/products/${product.slug}+${variantSlug}`);
    }
  };

  if (isLoading) return <FetchLoading />;
  if (isError) return <ErrorDialog onRetry={refetch} />;
  if (!product || !selectedVariant) return null;

  // Gallery berdasarkan warna aktif (gambar unik per warna)
  const galleryImages = Array.from(
    new Set(
      product.variants
        .filter((v) => v.options?.colors === selectedOptions.colors)
        .map((v) => v.imageUrl)
        .filter(Boolean)
    )
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
          selectedOptions={selectedOptions}
          onOptionChange={handleVariantOptionChange}
        />
      </div>
    </div>
  );
};

export default Product;
