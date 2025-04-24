import Logo from "@/components/ui/Logo";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import SearchInput from "@/components/header/SearchInput";
import SearchDropdown from "@/components/header/SearchDropdown";
import { useSearchProductsQuery } from "@/hooks/useProductsQuery";
import { useAuthStore } from "../../store/useAuthStore";
import { ShoppingCart, LogIn } from "lucide-react";

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
          <div className="flex items-center gap-4">
            <Logo />
          </div>

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
            {/* Keranjang hanya jika user login */}
            {user && (
              <button
                onClick={() => navigate("/cart")}
                className="relative text-muted-foreground hover:text-primary"
              >
                <ShoppingCart className="w-6 h-6" />
              </button>
            )}

            {/* Avatar / Login */}
            {user ? (
              <img
                src={user.profile?.avatar}
                alt={user.profile?.fullname}
                className="w-8 h-8 rounded-full object-cover cursor-pointer"
                onClick={() => navigate("/profile")}
              />
            ) : (
              <button
                onClick={handleLoginClick}
                className="flex items-center gap-2 text-sm border px-3 py-1.5 rounded-md hover:bg-gray-100 transition"
              >
                <LogIn className="w-4 h-4" />
                Login
              </button>
            )}
          </div>
        </div>
      </header>
    </div>
  );
};

export default Header;
