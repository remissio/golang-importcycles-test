# Importcycles Testing

This repo demonstrates how to handle "circular references" between two modules.

The packages `providera` and `providerb` each contain implementation logic for interfaces. They call each other which would result if implemented without interfaces in between them. The `interfaces` module contains interfaces for both of them.
