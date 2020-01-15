import React from "react";
import { View, Text, StyleSheet } from "react-native";

import color from "../../utils/colorPallet";

const DetailsScreen = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.textStyle}>
        Uhhh yeah, you son of a hard clicking bitch!
      </Text>
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
    fontSize: 30,
    fontWeight: "normal",
    color: color.brand.dark,
    textAlign: "center"
  }
});

export default DetailsScreen;
