import { useNavigate } from "react-router-dom";

const ProductCard = ({ product }) => {
  const navigate = useNavigate();

  return (
    <div
      onClick={() => navigate(`/products/${product.slug}`)}
      className="bg-white rounded-xl border shadow-sm hover:shadow-md transition duration-300 cursor-pointer group"
      key={product.id}
    >
      <div className="relative w-full aspect-square overflow-hidden rounded-t-xl">
        {/* Gambar utama */}
        <img
          src={product.images?.[0]}
          alt={product.name}
          className="absolute inset-0 w-full h-full object-cover transition-all duration-500 ease-in-out opacity-100 group-hover:opacity-0 scale-100 group-hover:scale-110"
        />
        {/* Gambar hover */}
        <img
          src="https://images.tokopedia.net/img/cache/500-square/VqbcmM/2023/3/26/7ff306d3-e514-441e-9b1d-f1c2f8a9aefc.jpg.webp?ect=4g"
          alt={product.name}
          className="absolute inset-0 w-full h-full object-cover transition-all duration-500 ease-in-out opacity-0 group-hover:opacity-100 scale-110 group-hover:scale-100"
        />
        {product.isFeatured && (
          <span className="absolute top-2 left-2 text-xs font-semibold bg-yellow-400 text-black px-2 py-1 rounded">
            ⭐ Featured
          </span>
        )}
      </div>

      <div className="p-4 space-y-1">
        <h2 className="text-base font-semibold text-gray-900">
          {product.name}
        </h2>
        <p className="text-xs text-muted-foreground">
          {product.subcategory?.name} — {product.category?.name}
        </p>
        <p className="text-sm text-gray-500 line-clamp-2">
          {product.description}
        </p>
        <p className="text-primary font-bold text-sm mt-2">${product.price}</p>
      </div>
    </div>
  );
};

export default ProductCard;
