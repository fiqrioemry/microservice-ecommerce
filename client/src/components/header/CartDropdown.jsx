// src/components/header/CartDropdown.jsx
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
} from "@/components/ui/dropdown-menu";
import { useEffect } from "react";
import { formatRupiah } from "@/lib/utils";
import { ShoppingCart } from "lucide-react";
import { useNavigate } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { useCartQuery } from "@/hooks/useCartQuery";
import { useCartStore } from "@/store/useCartStore";

const CartDropdown = () => {
  const navigate = useNavigate();
  const { setCart } = useCartStore();
  const { data: cart = [], isLoading } = useCartQuery();
  console.log(cart.items);
  useEffect(() => {
    if (cart.items) {
      setCart(cart.items);
    }
  }, [cart, setCart]);

  const { items, totalItems } = useCartStore();
  const hasItems = items.length > 0;

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <button className="relative text-muted-foreground hover:text-primary">
          <ShoppingCart className="w-6 h-6" />
          {totalItems > 0 && (
            <span className="absolute -top-2 -right-2 bg-red-500 text-white text-xs w-5 h-5 rounded-full flex items-center justify-center min-w-[20px] h-5 px-1 text-center leading-none">
              {totalItems}
            </span>
          )}
        </button>
      </DropdownMenuTrigger>

      <DropdownMenuContent
        align="end"
        className="w-[320px] p-0 shadow-lg rounded-xl overflow-hidden"
      >
        {isLoading ? (
          <div className="p-4 text-center text-sm text-muted-foreground">
            Loading...
          </div>
        ) : hasItems ? (
          <>
            <ul className="max-h-64 overflow-y-auto divide-y">
              {items.map((item) => (
                <li key={item.id} className="flex gap-3 p-4">
                  <img
                    src={item.imageUrl}
                    alt={item.name}
                    className="w-16 h-16 rounded-md object-cover border"
                  />
                  <div className="flex-1 text-sm">
                    <p className="font-medium line-clamp-2">{item.name}</p>
                    <p className="text-muted-foreground text-xs mt-1">
                      {item.quantity}x {formatRupiah(item.price)}
                    </p>
                  </div>
                </li>
              ))}
            </ul>
            <div className="p-4 border-t">
              <Button onClick={() => navigate("/cart")} className="w-full">
                View Cart
              </Button>
            </div>
          </>
        ) : (
          <div className="p-2 flex flex-col items-center justify-center text-center h-80">
            <img src="/empty-cart.webp" alt="Empty Cart" />
          </div>
        )}
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default CartDropdown;
