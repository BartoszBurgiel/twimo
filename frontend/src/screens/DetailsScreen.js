// Import core functionalities from react & react-native library
import React from "react";
import { View, Text, StyleSheet } from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define Component
const DetailsScreen = props => {
  return (
    <View style={styles.container}>
      <Text style={styles.textStyle}>built & designed by Bartosz & Louis</Text>
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
    fontSize: 25,
    fontWeight: "normal",
    color: color.brand.dark,
    textAlign: "center"
  }
});

// Export the above defined component
export default DetailsScreen;
