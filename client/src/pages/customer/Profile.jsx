import React from "react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { useProfileQuery } from "@/hooks/useUserQuery";

const Profile = () => {
  const { data: profile, isError, refetch, isLoading } = useProfileQuery();

  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog onRetry={refetch} />;

  console.log(profile);
  return (
    <div>
      Lorem ipsum, dolor sit amet consectetur adipisicing elit. Earum ut
      incidunt placeat natus similique beatae id aspernatur, minus rem est
      numquam quasi, consequuntur sed ea voluptatum quos, veniam vel nam.
    </div>
  );
};

export default Profile;
