import React from "react";
import { View, Text, Button, StyleSheet } from "react-native";

import color from "../../utils/colorPallet";

const HomeScreen = props => {
  return (
    <View style={styles.container}>
      <Text style={styles.textStyle}>Hello iOS...</Text>
      <Button
        title="Click me softly ..."
        onPress={() => props.navigation.navigate("List")}
      ></Button>
      <Button
        title="... or really hard!"
        onPress={() => props.navigation.navigate("Details")}
      ></Button>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.white,
    alignItems: "center",
    justifyContent: "center"
  },
  textStyle: {
    fontSize: 50,
    fontWeight: "bold",
    color: color.brand.normal
  }
});

export default HomeScreen;
