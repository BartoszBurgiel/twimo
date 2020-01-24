// Import core functionality dependencies from react, react-native and expo
import React from "react";
import { View, Text, StyleSheet, Image, TouchableOpacity } from "react-native";
import { Ionicons } from "@expo/vector-icons";

// Import custom components
import RatingStars from "./RatingStars";

// Import the color utility
import color from "../../utils/colorPallet";

// Define LocationCard component
// LocationCard receives locationdata and the navigation object as properties from the parent component
const LocationCard = ({ data, navigation }) => {
  return (
    <View style={styles.card}>
      <View style={styles.pictureWrapper}>
        <Image style={styles.image} source={{ uri: data.picture }}></Image>
      </View>
      <View style={styles.textWrapper}>
        <Text style={styles.cardHeader}>{data.name}</Text>
        <Text style={styles.cardSubHeader}>{data.location}</Text>
        <RatingStars colorFill={color.brand.dark} rating={data.rating} />
        <Text style={styles.cardMetrics}>
          <Ionicons name="logo-euro" size={20} color={color.brand.dark} />{" "}
          {data.pricing} / 5
        </Text>
        <TouchableOpacity
          style={styles.button}
          onPress={() =>
            navigation.navigate("LocationDetails", {
              locationData: data
            })
          }
        >
          <Text style={styles.buttonText}>Mehr Infos</Text>
        </TouchableOpacity>
      </View>
    </View>
  );
};

// Define the stylesheet
const styles = StyleSheet.create({
  card: {
    backgroundColor: color.white,
    width: 370,
    height: 200,
    padding: 10,
    marginVertical: 10,
    elevation: 5,
    shadowColor: color.gray.darker,
    shadowOpacity: 0.25,
    borderRadius: 12,
    shadowOffset: {
      width: 0,
      height: 5
    },
    shadowRadius: 13,
    flexDirection: "row"
  },
  image: {
    height: 180,
    width: 180,
    borderRadius: 12
  },
  textWrapper: {
    justifyContent: "space-around",
    paddingLeft: 10,
    width: 170
  },
  cardHeader: {
    fontWeight: "900",
    fontSize: 25,
    color: color.brand.normal
  },
  cardSubHeader: {
    fontWeight: "300",
    textTransform: "uppercase",
    fontSize: 12
  },
  cardMetrics: {
    fontWeight: "600"
  },
  button: {
    backgroundColor: color.brand.light,
    height: 30,
    width: 120,
    borderRadius: 12,
    justifyContent: "center",
    alignItems: "center",
    marginTop: 20
  },
  buttonText: {
    textTransform: "uppercase",
    color: color.brand.dark,
    fontWeight: "800"
  }
});

// Export the Card component to be used across the App
export default LocationCard;
