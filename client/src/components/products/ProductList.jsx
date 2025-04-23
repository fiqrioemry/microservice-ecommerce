import { Link } from "react-router-dom";

const ProductList = ({ product }) => {
  const image =
    product.images?.[0] || "https://placehold.co/100x100?text=No+Image";

  return (
    <div className="flex gap-4 border-b py-4" key={product.id}>
      <img
        src={image}
        alt={product.name}
        className="w-24 h-24 object-cover rounded"
      />
      <div className="flex-1">
        <h3 className="text-lg font-semibold">{product.name}</h3>
        <p className="text-sm text-muted-foreground">
          {product.description?.slice(0, 80)}...
        </p>
        <p className="mt-1 text-primary font-semibold">
          Rp {product.price?.toLocaleString()}
        </p>
        <p className="text-sm text-muted-foreground mt-1">
          {product.category?.name} / {product.subcategory?.name}
        </p>
        <div className="mt-2">
          <Link
            to={`/products/${product.slug}`}
            className="inline-block text-sm bg-primary text-white px-3 py-1 rounded"
          >
            View Detail
          </Link>
        </div>
      </div>
    </div>
  );
};

export default ProductList;
