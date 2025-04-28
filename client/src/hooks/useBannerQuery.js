import banner from "@/services/banners";
import { useQuery } from "@tanstack/react-query";

export const useGetAllBannersQuery = () =>
  useQuery({
    queryKey: ["banners"],
    queryFn: banner.getAllBanners,
  });

export const useGetBannerByPositionQuery = (position) =>
  useQuery({
    queryKey: ["banners", position],
    queryFn: () => banner.getBannersByPosition(position),
    enabled: !!position,
  });
