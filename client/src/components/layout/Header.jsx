import { useAuthStore } from "@/store/useAuthStore";

const Header = () => {
  const { user } = useAuthStore();

  return (
    <div className="h-14">
      <header className="fixed w-full z-50 bg-background p-2 border-b">
        <div className="flex items-center justify-between container mx-auto gap-4">
          {/* Website Logo */}
          <div className="hidden md:flex px-2">
            <h4>LOGO</h4>
          </div>
        </div>
      </header>
    </div>
  );
};

export default Header;
