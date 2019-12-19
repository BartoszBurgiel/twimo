import React, { useState } from "react";
import { StyleSheet, Text, View, Image } from "react-native";
import { Toolbar } from "react-native-material-ui";

function Greeting(props) {
  return (
    <View style={{ alignItems: "center" }}>
      <Text>Hello {props.name}</Text>
    </View>
  );
}

export default function LotsOfGreetings() {
  return (
    <View>
      <Toolbar
        leftElement="menu"
        centerElement="Searchable"
        searchable={{
          autoFocus: true,
          placeholder: "Search"
        }}
        rightElement={{
          menu: {
            icon: "more-vert",
            labels: ["item 1", "item 2"]
          }
        }}
        onRightElementPress={label => {
          console.log(label);
        }}
      />
    </View>
  );
}
