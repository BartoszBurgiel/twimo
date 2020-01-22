// Import core functionality dependencies
import React from "react";
import { View, Text, StyleSheet, Image } from "react-native";

// Import color utility for consistent styling
import color from "../../utils/colorPallet";

// Define Component
const LocationDetailsScreen = () => {
  return (
    <View style={styles.container}>
      <Text>Some Picture</Text>
      <Text>Kind of location</Text>
      <Text>Name</Text>
      <Text>Rating</Text>
      <Text>SupplyIcons</Text>
      <Text>Separator</Text>
      <Text>UserRatings</Text>
    </View>
  );
};

// Define styles
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: color.white,
    alignItems: "center",
    justifyContent: "center"
  }
});

// Export the above defined component
export default LocationDetailsScreen;
