/* eslint-disable react/prop-types */
import { useEffect } from "react";
import Loading from "@/components/ui/Loading";
import { useAuthStore } from "@/store/useAuthStore";

const AuthProvider = ({ children }) => {
  const { user, authCheck, checkingAuth } = useAuthStore();

  useEffect(() => {
    authCheck();
  }, [authCheck]);

  if (checkingAuth) return <Loading />;

  return <>{children}</>;
};

export default AuthProvider;
