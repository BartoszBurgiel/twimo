// Import core functionalities from react & react-native library
import React from "react";
import { View, Text, StyleSheet } from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define Component
const DetailsScreen = () => {
  return (
    <View style={styles.container}>
      <Text style={styles.textStyle}>
        Uhhh yeah, you son of a hard clicking bitch!
      </Text>
    </View>
  );
};

// Define the styles
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

// Export the above defined component
export default DetailsScreen;
