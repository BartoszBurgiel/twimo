import React, { useState } from "react";
import { StyleSheet, Text, View, Image } from "react-native";
import ListOfCards from "./app/components/ListOfCards";

export default function Greeting(props) {
  return (
    <View style={{ alignItems: "center" }}>
      <ListOfCards />
    </View>
  );
}
