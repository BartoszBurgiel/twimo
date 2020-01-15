import React from "react";
import { View, Text, Button, StyleSheet } from "react-native";

import color from "../../utils/colorPallet";

const HomeScreen = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.textStyle}>Hello iOS!</Text>
      <Button title="trolololol"></Button>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.brand.light,
    alignItems: "center",
    justifyContent: "center"
  },
  textStyle: {
    fontSize: 50,
    fontWeight: "bold",
    color: color.brand.dark
  }
});

export default HomeScreen;