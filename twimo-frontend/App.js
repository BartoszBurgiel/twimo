import React, { useState } from "react";
import { StyleSheet, Text, View, Image } from "react-native";

function Greeting(props) {
  return (
    <View style={{ alignItems: "center" }}>
      <Text>Hello {props.name}</Text>
    </View>
  );
}

export default function LotsOfGreetings() {
  return (
    <View style={{ alignItems: "center", top: 50 }}>
      <Greeting name="Louis" />
      <Greeting name="Bartosz" />
      <Greeting name="Whoever" />
    </View>
  );
}
