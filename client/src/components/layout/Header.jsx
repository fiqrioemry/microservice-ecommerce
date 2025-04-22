import { useState, useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { useSearchProductsQuery } from "@/hooks/useProductsQuery";
import { useAuthStore } from "@/store/useAuthStore";

const Header = () => {
  const { user } = useAuthStore();
  const navigate = useNavigate();
  const [search, setSearch] = useState("");
  const [debouncedSearch, setDebouncedSearch] = useState("");
  const inputRef = useRef(null);

  useEffect(() => {
    const handler = setTimeout(() => setDebouncedSearch(search), 500); // debounce 500ms
    return () => clearTimeout(handler);
  }, [search]);

  const { data, isFetching } = useSearchProductsQuery(
    debouncedSearch ? { q: debouncedSearch, limit: 5 } : null
  );

  const handleKeyDown = (e) => {
    if (e.key === "Enter") {
      navigate(`/product?q=${search}`);
      setSearch("");
    }
  };

  const handleResultClick = (slug) => {
    navigate(`/products/${slug}`);
    setSearch("");
  };

  return (
    <div className="h-14 relative z-50">
      <header className="fixed w-full bg-background p-2 border-b shadow-sm">
        <div className="flex items-center justify-between container mx-auto gap-4 relative">
          {/* Logo */}
          <div className="hidden md:flex px-2">
            <h4>LOGO</h4>
          </div>

          {/* Search Input */}
          <div className="relative w-full max-w-md">
            <input
              ref={inputRef}
              type="text"
              placeholder="Search products..."
              value={search}
              onChange={(e) => setSearch(e.target.value)}
              onKeyDown={handleKeyDown}
              className="w-full border border-gray-300 rounded-md px-3 py-2 text-sm focus:outline-none focus:ring-2 focus:ring-primary"
            />

            {/* Dropdown Results */}
            {search && (
              <div className="absolute top-full left-0 right-0 bg-white border border-gray-200 shadow-md rounded-md mt-1 max-h-60 overflow-y-auto z-50">
                {isFetching ? (
                  <div className="p-3 text-sm text-gray-500">Loading...</div>
                ) : data?.results?.length ? (
                  data.results.map((item) => (
                    <div
                      key={item.id}
                      onClick={() => handleResultClick(item.slug)}
                      className="p-3 text-sm hover:bg-gray-100 cursor-pointer"
                    >
                      {item.name}
                    </div>
                  ))
                ) : (
                  <div className="p-3 text-sm text-gray-500">No result</div>
                )}
              </div>
            )}
          </div>
        </div>
      </header>
    </div>
  );
};

export default Header;
