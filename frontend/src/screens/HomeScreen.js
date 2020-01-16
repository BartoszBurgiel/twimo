// Import core functionalities from react & react-native library
import React from "react";
import { View, Text, Button, StyleSheet } from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define Component
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

// Define the styles
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

// Export the above defined component
export default HomeScreen;
