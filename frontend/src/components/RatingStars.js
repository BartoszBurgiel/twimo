// Import core dependencies from react, react native and expo
import React from "react";
import { Text } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Define RatingScale component
// Receives rating, colorfill and (optional) iconSize as properties from parent component
const RatingStars = ({ rating, colorFill, iconSize }) => {
  // Set up targetArray
  let arrayOfIcons = [];
  // Loop five times
  for (let i = 0; i < 5; i++) {
    if (i < rating) {
      // Print (rating) filled stars
      arrayOfIcons.push(
        <Ionicons
          name="md-star"
          key={i}
          size={iconSize ? iconSize : 25}
          color={colorFill}
        />
      );
    } else {
      // Print (5-rating) empty stars
      arrayOfIcons.push(
        <Ionicons
          name="md-star-outline"
          key={i}
          size={iconSize ? iconSize : 25}
          color={colorFill}
        />
      );
    }
  }
  // Return text element containing the array
  return <Text>{arrayOfIcons}</Text>;
};

// Export the component to be usable across the app
export default RatingStars;
