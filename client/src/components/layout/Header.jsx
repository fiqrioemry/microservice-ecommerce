import Logo from "@/components/ui/Logo";
import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import SearchInput from "@/components/header/SearchInput";
import SearchDropdown from "@/components/header/SearchDropdown";
import { useSearchProductsQuery } from "@/hooks/useProductsQuery";

const Header = () => {
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
    navigate(`/products/${slug}`);
    setSearch("");
  };

  return (
    <div className="h-14 relative z-50">
      <header className="fixed w-full bg-background p-2 border-b shadow-sm">
        <div className="flex items-center justify-between container mx-auto gap-4 relative">
          <div className="hidden md:flex px-2">
            <Logo />
          </div>

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
        </div>
      </header>
    </div>
  );
};

export default Header;
