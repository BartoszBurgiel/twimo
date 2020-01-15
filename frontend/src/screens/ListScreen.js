import React from "react";
import { View, Text, StyleSheet } from "react-native";

import color from "../../utils/colorPallet";

const ListScreen = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.softTextStyles}>
        That didn't even count as a click...
      </Text>
      <Text style={styles.hardTextStyles}>Go harder next time!</Text>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.black,
    alignItems: "center",
    justifyContent: "center"
  },
  softTextStyles: {
    fontSize: 30,
    fontWeight: "normal",
    color: color.brand.light,
    textAlign: "center"
  },
  hardTextStyles: {
    fontSize: 30,
    fontWeight: "bold",
    color: color.brand.normal,
    textAlign: "center"
  }
});

export default ListScreen;
