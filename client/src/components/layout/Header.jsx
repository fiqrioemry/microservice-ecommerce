import { useState, useEffect } from "react";
import Button from "@/components/ui/button";
import { Link, useNavigate } from "react-router-dom";
import { ShoppingCart, LogIn } from "lucide-react";
import { useAuthStore } from "@/store/useAuthStore";
import SearchInput from "@/components/header/SearchInput";
import SearchDropdown from "@/components/header/SearchDropdown";
import { useSearchProductsQuery } from "@/hooks/useProductsQuery";

const Header = () => {
  const { user } = useAuthStore();
  const navigate = useNavigate();
  const [search, setSearch] = useState("");
  const [debouncedSearch, setDebouncedSearch] = useState("");

  useEffect(() => {
    const handler = setTimeout(() => setDebouncedSearch(search), 500);
    return () => clearTimeout(handler);
  }, [search]);

  const { data, isFetching } = useSearchProductsQuery(
    debouncedSearch ? { q: debouncedSearch, limit: 5 } : null
  );

  const handleResultClick = (slug) => {
    console.log(slug);
    navigate(`/products/${slug}`);
    setSearch("");
  };

  const handleLoginClick = () => navigate("/signin");

  return (
    <div className="h-14 relative z-50">
      <header className="fixed w-full bg-white p-2 border-b shadow-sm">
        <div className="flex items-center justify-between container mx-auto gap-4">
          {/* Logo */}

          <Link to="/">
            <h2 className="text-xl font-bold text-primary">Ecommerce</h2>
          </Link>

          {/* Search */}
          <div className="relative w-full max-w-md">
            <SearchInput
              value={search}
              onChange={setSearch}
              onKeyDown={(e) => {
                if (e.key === "Enter") {
                  navigate(`/products?q=${search}`);
                  setSearch("");
                }
              }}
              onClear={() => setSearch("")}
            />
            <SearchDropdown
              search={search}
              isLoading={isFetching}
              results={data?.results}
              onClick={handleResultClick}
            />
          </div>

          {/* Right section */}
          <div className="flex items-center gap-4">
            {/* Shopping cart dropdown*/}
            {user && (
              <button
                onClick={() => navigate("/cart")}
                className="relative text-muted-foreground hover:text-primary"
              >
                <ShoppingCart className="w-6 h-6" />
              </button>
            )}

            {/* User Dropdown Avatar & Login */}
            {user ? (
              <img
                src={user.profile?.avatar}
                alt={user.profile?.fullname}
                onClick={() => navigate("/profile")}
                className="w-8 h-8 rounded-full object-cover cursor-pointer"
              />
            ) : (
              <Button onClick={handleLoginClick}>
                <LogIn className="w-4 h-4" />
                Login
              </Button>
            )}
          </div>
        </div>
      </header>
    </div>
  );
};

export default Header;
