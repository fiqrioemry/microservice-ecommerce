import React from "react";
import { Outlet } from "react-router-dom";
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs";

const CustomerLayout = () => {
  return (
    <Tabs defaultValue="account" className="section">
      <TabsList>
        <TabsTrigger value="profile">profile</TabsTrigger>
        <TabsTrigger value="address">address</TabsTrigger>
        <TabsTrigger value="orders">orders</TabsTrigger>
      </TabsList>
      <Outlet />
    </Tabs>
  );
};

export default CustomerLayout;
