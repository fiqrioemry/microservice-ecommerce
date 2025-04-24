import React, { useEffect, useState } from "react";
import clsx from "clsx";
import { useCartStore } from "@/store/useCartStore";
import { Button } from "@/components/ui/button";
import { Minus, Plus } from "lucide-react";

const ProductVariantSelector = ({
  product,
  selectedVariant,
  onSelectVariant,
}) => {
  const [quantity, setQuantity] = useState(1);
  const [activeOptions, setActiveOptions] = useState(
    selectedVariant.options || {}
  );
  const addItem = useCartStore((state) => state.addItem);

  useEffect(() => {
    setActiveOptions(selectedVariant.options || {});
  }, [selectedVariant]);

  const handleOptionClick = (key, value) => {
    const updatedOptions = { ...activeOptions, [key]: value };
    setActiveOptions(updatedOptions);

    const matchedVariant = product.variants.find((v) =>
      Object.entries(updatedOptions).every(
        ([k, vOpt]) => v.options?.[k] === vOpt
      )
    );

    if (matchedVariant) {
      onSelectVariant(matchedVariant);
    }
  };

  const handleAddToCart = async () => {
    if (quantity > selectedVariant.stock) return;
    await addItem({
      productId: product.id,
      variantId: selectedVariant.id,
      quantity,
    });
  };

  const optionKeys = Object.keys(selectedVariant.options || {});

  return (
    <>
      {optionKeys.map((key) => {
        const uniqueOptionValues = [
          ...new Set(product.variants.map((v) => v.options?.[key])),
        ];

        return (
          <div key={key} className="mt-4">
            <p className="text-sm font-medium text-muted-foreground mb-1">
              Pilih {key}
            </p>
            <div className="flex gap-2 flex-wrap">
              {uniqueOptionValues.map((optionValue) => {
                const isActive = activeOptions?.[key] === optionValue;

                return (
                  <button
                    key={key + optionValue}
                    onClick={() => handleOptionClick(key, optionValue)}
                    className={clsx(
                      "px-3 py-1 border rounded text-sm",
                      isActive
                        ? "bg-primary text-white border-primary"
                        : "bg-white text-gray-700 border-gray-300 hover:border-primary"
                    )}
                  >
                    {optionValue}
                  </button>
                );
              })}
            </div>
          </div>
        );
      })}

      {/* Quantity Selector */}
      <div className="flex items-center gap-3 mt-6">
        <p className="text-sm font-medium text-muted-foreground">Jumlah:</p>
        <div className="flex items-center border rounded-md px-2">
          <button
            onClick={() => setQuantity(Math.max(1, quantity - 1))}
            className="p-1 text-gray-600 hover:text-primary"
          >
            <Minus className="w-4 h-4" />
          </button>
          <span className="px-3 text-sm font-medium">{quantity}</span>
          <button
            onClick={() =>
              setQuantity(Math.min(quantity + 1, selectedVariant.stock))
            }
            className="p-1 text-gray-600 hover:text-primary"
          >
            <Plus className="w-4 h-4" />
          </button>
        </div>
        <span className="text-xs text-muted-foreground">
          Stok: {selectedVariant.stock}
        </span>
      </div>

      {/* Add to Cart Button */}
      <div className="mt-4">
        <Button
          disabled={
            selectedVariant.stock === 0 || quantity > selectedVariant.stock
          }
          onClick={handleAddToCart}
          className="w-full"
        >
          Tambah ke Keranjang
        </Button>
      </div>
    </>
  );
};

export default ProductVariantSelector;
